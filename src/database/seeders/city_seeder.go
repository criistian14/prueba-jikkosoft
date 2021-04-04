package seeders

import (
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	cityRepository "github.com/criistian14/prueba-jikkosoft/src/modules/cities/repository"
	cityUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/cities/usecase"
	"time"
)

func CitySeeder() error {
	cityRepo := cityRepository.NewCityRepository()
	cityUsecase := cityUseCase.NewCityUsecase(cityRepo)

	// City 1
	var countryID uint = 1
	name := "Cali"

	_, err := cityUsecase.SaveCity(cityDomain.City{
		Name:      &name,
		CountryID: &countryID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	// City 2
	name = "Bogota"

	_, err = cityUsecase.SaveCity(cityDomain.City{
		Name:      &name,
		CountryID: &countryID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
