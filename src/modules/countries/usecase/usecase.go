package usecase

import (
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
	countryRepository "github.com/criistian14/prueba-jikkosoft/src/modules/countries/repository"
)

type CountryUsecase interface {
	GetCountry(country countryDomain.Country, isDeleted bool) (countryDomain.Country, error)
	SaveCountry(country countryDomain.Country) (countryDomain.Country, error)
}

type countryUsecase struct {
	repo countryRepository.CountryRepository
}

// * Create and return new instance
func NewCountryUsecase(repo countryRepository.CountryRepository) *countryUsecase {
	return &countryUsecase{
		repo: repo,
	}
}
