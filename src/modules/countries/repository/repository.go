package repository

import countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"

type CountryRepository interface {
	FindCountry(country countryDomain.Country, isDeleted bool) (countryDomain.Country, error)
	SaveCountry(country countryDomain.Country) (countryDomain.Country, error)
}

type countryRepository struct {
}

// * Create and return new instance
func NewCountryRepository() *countryRepository {
	return &countryRepository{}
}
