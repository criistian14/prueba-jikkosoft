package usecase

import (
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
)

func (usecase publicServiceUsecase) GetPublicService(publicService publicServiceDomain.PublicService, isDeleted bool) (publicServiceDomain.PublicService, error) {
	return usecase.repo.FindPublicService(publicService, isDeleted)
}
