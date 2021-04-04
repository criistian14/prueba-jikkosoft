package repository

import (
	"errors"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
	"gorm.io/gorm"
)

func (repo countryRepository) FindCountry(countryToSearch countryDomain.Country, isDeleted bool) (countryDomain.Country, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return countryToSearch, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	var country countryDomain.Country
	if isDeleted {
		db = db.Unscoped()
	}

	result := db.Where(&countryToSearch).First(&country)
	// Controlled error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return countryDomain.Country{}, projectErrors.ErrCountryNotFound
		}

		// Unknown error
		return countryDomain.Country{}, result.Error
	}

	return country, nil
}
