package repository

import (
	projectConditionals "github.com/criistian14/prueba-jikkosoft/src/conditionals"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
)

func (repo cityRepository) SaveCity(city cityDomain.City) (cityDomain.City, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return city, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	result := db.Create(&city)
	// Controlled error
	if result.Error != nil {
		// Duplicate ID
		if projectConditionals.IDIsDuplicated(result.Error) {
			return city, projectErrors.ErrCityDuplicateID
		}

		// Unknown error
		return city, result.Error
	}

	return city, nil
}
