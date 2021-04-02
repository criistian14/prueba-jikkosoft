package seeders

import (
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	publicServiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/repository"
	publicServiceUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/usecase"
	"time"
)

func PublicServiceSeeder() error {
	publicServiceRepo := publicServiceRepository.NewPublicServiceRepository()
	publicServiceUsecase := publicServiceUseCase.NewPublicServiceUsecase(publicServiceRepo)

	// Company 1
	company := "Company Water"
	typePublicService := publicServiceDomain.PublicWaterService
	email := "water@water.com"

	_, err := publicServiceUsecase.SavePublicService(publicServiceDomain.PublicService{
		Company:   &company,
		Type:      &typePublicService,
		Email:     &email,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	// Company 2
	company = "Company Electric"
	typePublicService = publicServiceDomain.PublicElectricService
	email = "electric@electric.com"

	_, err = publicServiceUsecase.SavePublicService(publicServiceDomain.PublicService{
		Company:   &company,
		Type:      &typePublicService,
		Email:     &email,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
