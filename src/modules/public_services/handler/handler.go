package handler

import (
	publicServiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/repository"
	publicServiceUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/usecase"
)

type PublicServiceHandler interface {
}

type publicServiceHandler struct {
	publicServiceUsecase publicServiceUseCase.PublicServiceUsecase
}

// * Create and return new instance
func NewPublicServiceHandler() *publicServiceHandler {
	publicServiceRepo := publicServiceRepository.NewPublicServiceRepository()
	publicServiceUsecase := publicServiceUseCase.NewPublicServiceUsecase(publicServiceRepo)

	return &publicServiceHandler{
		publicServiceUsecase: publicServiceUsecase,
	}
}
