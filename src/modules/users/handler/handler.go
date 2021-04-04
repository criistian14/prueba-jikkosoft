package handler

import (
	cityRepository "github.com/criistian14/prueba-jikkosoft/src/modules/cities/repository"
	cityUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/cities/usecase"
	publicServiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/repository"
	publicServiceUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/usecase"
	userRepository "github.com/criistian14/prueba-jikkosoft/src/modules/users/repository"
	userUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/users/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	FindUserHandler(context *fiber.Ctx) error
}

type userHandler struct {
	userUsecase          userUseCase.UserUsecase
	cityUsecase          cityUseCase.CityUsecase
	publicServiceUsecase publicServiceUseCase.PublicServiceUsecase
}

// * Create and return new instance
func NewUserHandler() *userHandler {
	userRepo := userRepository.NewUserRepository()
	userUsecase := userUseCase.NewUserUsecase(userRepo)

	cityRepo := cityRepository.NewCityRepository()
	cityUsecase := cityUseCase.NewCityUsecase(cityRepo)

	publicServiceRepo := publicServiceRepository.NewPublicServiceRepository()
	publicServiceUsecase := publicServiceUseCase.NewPublicServiceUsecase(publicServiceRepo)

	return &userHandler{
		userUsecase:          userUsecase,
		cityUsecase:          cityUsecase,
		publicServiceUsecase: publicServiceUsecase,
	}
}
