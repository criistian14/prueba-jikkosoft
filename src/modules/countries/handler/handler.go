package handler

import (
	countryRepository "github.com/criistian14/prueba-jikkosoft/src/modules/countries/repository"
	countryUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/countries/usecase"
)

type CountryHandler interface {
}

type countryHandler struct {
	countryUsecase countryUseCase.CountryUsecase
}

// * Create and return new instance
func NewCountryHandler() *countryHandler {
	countryRepo := countryRepository.NewCountryRepository()
	countryUsecase := countryUseCase.NewCountryUsecase(countryRepo)

	return &countryHandler{
		countryUsecase: countryUsecase,
	}
}
