package handler

import (
	numberDomain "github.com/criistian14/prueba-jikkosoft/src/modules/numbers/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (handler *numberHandler) SortArrayHandler(context *fiber.Ctx) error {
	data := new(numberDomain.RequestData)
	if err := context.BodyParser(data); err != nil {
		return context.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err := data.Validate(numberDomain.FieldsToValidate{
		Numbers: true,
	})
	if err != nil {
		return context.Status(http.StatusBadRequest).SendString(err.Error())
	}

	numbersUnsorted := data.Numbers
	numbersSorted, err := handler.numberUsecase.SortArrayNumbers(data.Numbers)
	if err != nil {
		return context.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"unsorted": numbersUnsorted,
		"sorted":   numbersSorted,
	})
}
