package handler

import (
	"errors"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (handler *userHandler) FindUserHandler(context *fiber.Ctx) error {
	var userToSearch userDomain.User
	response := make(map[string]interface{})

	idUser, err := strconv.ParseUint(context.Params("id"), 10, 32)
	if err != nil {
		return context.Status(http.StatusBadRequest).SendString(err.Error())
	}
	userToSearch.ID = uint(idUser)

	userFound, err := handler.userUsecase.GetUser(userToSearch, false)
	if err != nil {
		if errors.Is(err, projectErrors.ErrUserNotFound) {
			return context.Status(http.StatusNotFound).SendString(err.Error())
		}

		return context.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	response["user"] = userFound

	// Get city
	if userFound.CityID != nil {
		cityFound, err := handler.cityUsecase.GetCity(cityDomain.City{
			ID: userFound.ID,
		}, false)
		if err != nil {
			if errors.Is(err, projectErrors.ErrCityNotFound) {
				return context.Status(http.StatusNotFound).SendString(err.Error())
			}

			return context.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		response["city"] = cityFound
	}

	// Get All Public Services
	if len(userFound.Invoices) > 0 {
		var publicServices []publicServiceDomain.PublicService
		checkRepeat := make(map[uint]bool)

		for _, invoice := range userFound.Invoices {
			// Check that the same public service is not consulted again.
			if checkRepeat[*invoice.PublicServiceID] {
				continue
			}

			publicServiceFound, err := handler.publicServiceUsecase.GetPublicService(publicServiceDomain.PublicService{
				ID: *invoice.PublicServiceID,
			}, false)
			if err != nil {
				if errors.Is(err, projectErrors.ErrInvoiceNotFound) {
					return context.Status(http.StatusNotFound).SendString(err.Error())
				}

				return context.Status(http.StatusInternalServerError).SendString(err.Error())
			}

			checkRepeat[*invoice.PublicServiceID] = true
			publicServices = append(publicServices, publicServiceFound)
		}

		response["public_services"] = publicServices
	}

	return context.Status(http.StatusOK).JSON(response)
}
