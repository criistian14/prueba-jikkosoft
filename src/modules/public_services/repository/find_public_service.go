package repository

import (
	"errors"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	"gorm.io/gorm"
)

func (repo publicServiceRepository) FindPublicService(publicServiceToSearch publicServiceDomain.PublicService, isDeleted bool) (publicServiceDomain.PublicService, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return publicServiceToSearch, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	var publicService publicServiceDomain.PublicService
	if isDeleted {
		db = db.Unscoped()
	}

	result := db.Where(&publicServiceToSearch).First(&publicService)
	// Controlled error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return publicServiceDomain.PublicService{}, projectErrors.ErrPublicServiceNotFound
		}

		// Unknown error
		return publicServiceDomain.PublicService{}, result.Error
	}

	return publicService, nil
}
