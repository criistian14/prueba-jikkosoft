package usecase

import (
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	cityRepository "github.com/criistian14/prueba-jikkosoft/src/modules/cities/repository"
)

type CityUsecase interface {
	GetCity(city cityDomain.City, isDeleted bool) (cityDomain.City, error)
	SaveCity(city cityDomain.City) (cityDomain.City, error)
}

type cityUsecase struct {
	repo cityRepository.CityRepository
}

// * Create and return new instance
func NewCityUsecase(repo cityRepository.CityRepository) *cityUsecase {
	return &cityUsecase{
		repo: repo,
	}
}
