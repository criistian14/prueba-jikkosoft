package usecase

import (
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
)

func (usecase cityUsecase) GetCity(city cityDomain.City, isDeleted bool) (cityDomain.City, error) {
	return usecase.repo.FindCity(city, isDeleted)
}
