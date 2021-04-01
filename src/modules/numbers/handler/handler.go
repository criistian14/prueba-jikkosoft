package handler

import (
	"github.com/gofiber/fiber/v2"

	numberUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/numbers/usecase"
)

type NumberHandler interface {
	SortArrayHandler(context *fiber.Ctx) error
}

type numberHandler struct {
	numberUsecase numberUseCase.NumberUsecase
}

// * Create and return new instance
func NewNumberHandler() *numberHandler {
	numberUsecase := numberUseCase.NewNumberUsecase()

	return &numberHandler{
		numberUsecase: numberUsecase,
	}
}
