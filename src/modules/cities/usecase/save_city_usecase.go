package usecase

import (
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
)

func (usecase cityUsecase) SaveCity(city cityDomain.City) (cityDomain.City, error) {
	return usecase.repo.SaveCity(city)
}
