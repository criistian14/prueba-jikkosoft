package repository

import (
	projectConditionals "github.com/criistian14/prueba-jikkosoft/src/conditionals"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
)

func (repo countryRepository) SaveCountry(country countryDomain.Country) (countryDomain.Country, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return country, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	result := db.Create(&country)
	// Controlled error
	if result.Error != nil {
		// Duplicate ID
		if projectConditionals.IDIsDuplicated(result.Error) {
			return country, projectErrors.ErrCountryDuplicateID
		}

		// Unknown error
		return country, result.Error
	}

	return country, nil
}
