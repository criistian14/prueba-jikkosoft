package handler

import (
	cityRepository "github.com/criistian14/prueba-jikkosoft/src/modules/cities/repository"
	cityUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/cities/usecase"
)

type CityHandler interface {
}

type cityHandler struct {
	cityUsecase cityUseCase.CityUsecase
}

// * Create and return new instance
func NewCityHandler() *cityHandler {
	cityRepo := cityRepository.NewCityRepository()
	cityUsecase := cityUseCase.NewCityUsecase(cityRepo)

	return &cityHandler{
		cityUsecase: cityUsecase,
	}
}
