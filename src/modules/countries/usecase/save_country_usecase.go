package usecase

import (
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
)

func (usecase countryUsecase) SaveCountry(country countryDomain.Country) (countryDomain.Country, error) {
	return usecase.repo.SaveCountry(country)
}
