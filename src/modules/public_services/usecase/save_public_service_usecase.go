package usecase

import (
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
)

func (usecase publicServiceUsecase) SavePublicService(publicService publicServiceDomain.PublicService) (publicServiceDomain.PublicService, error) {
	return usecase.repo.SavePublicService(publicService)
}
