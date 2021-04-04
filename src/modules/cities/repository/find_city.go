package repository

import (
	"errors"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	"gorm.io/gorm"
)

func (repo cityRepository) FindCity(cityToSearch cityDomain.City, isDeleted bool) (cityDomain.City, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return cityToSearch, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	var city cityDomain.City
	if isDeleted {
		db = db.Unscoped()
	}

	result := db.Where(&cityToSearch).First(&city)
	// Controlled error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return cityDomain.City{}, projectErrors.ErrCityNotFound
		}

		// Unknown error
		return cityDomain.City{}, result.Error
	}

	return city, nil
}
