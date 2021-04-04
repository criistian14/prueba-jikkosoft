package usecase

import (
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
)

func (usecase countryUsecase) GetCountry(country countryDomain.Country, isDeleted bool) (countryDomain.Country, error) {
	return usecase.repo.FindCountry(country, isDeleted)
}
