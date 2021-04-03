package seeders

import (
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
	countryRepository "github.com/criistian14/prueba-jikkosoft/src/modules/countries/repository"
	countryUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/countries/usecase"
	"time"
)

func CountrySeeder() error {
	countryRepo := countryRepository.NewCountryRepository()
	countryUsecase := countryUseCase.NewCountryUsecase(countryRepo)

	// Country 1
	name := "Colombia"

	_, err := countryUsecase.SaveCountry(countryDomain.Country{
		Name:      &name,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	// Country 2
	name = "United States"

	_, err = countryUsecase.SaveCountry(countryDomain.Country{
		Name:      &name,
		CreatedAt: time.Now().AddDate(0, 0, -24),
	})
	if err != nil {
		return err
	}

	return nil
}
