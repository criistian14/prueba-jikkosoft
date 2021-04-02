package repository

import publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"

type PublicServiceRepository interface {
	FindPublicService(publicService publicServiceDomain.PublicService, isDeleted bool) (publicServiceDomain.PublicService, error)
	SavePublicService(publicService publicServiceDomain.PublicService) (publicServiceDomain.PublicService, error)
}

type publicServiceRepository struct {
}

// * Create and return new instance
func NewPublicServiceRepository() *publicServiceRepository {
	return &publicServiceRepository{}
}
