package usecase

import (
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	publicServiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/repository"
)

type PublicServiceUsecase interface {
	GetPublicService(publicService publicServiceDomain.PublicService, isDeleted bool) (publicServiceDomain.PublicService, error)
	SavePublicService(publicService publicServiceDomain.PublicService) (publicServiceDomain.PublicService, error)
}

type publicServiceUsecase struct {
	repo publicServiceRepository.PublicServiceRepository
}

// * Create and return new instance
func NewPublicServiceUsecase(repo publicServiceRepository.PublicServiceRepository) *publicServiceUsecase {
	return &publicServiceUsecase{
		repo: repo,
	}
}
