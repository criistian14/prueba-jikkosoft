package repository

import (
	"errors"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
	"gorm.io/gorm"
)

func (repo userRepository) FindUser(userToSearch userDomain.User, isDeleted bool) (userDomain.User, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return userToSearch, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	var user userDomain.User
	if isDeleted {
		db = db.Unscoped()
	}

	result := db.Where(&userToSearch).Preload("Invoices").Preload("Inquiries").First(&user)
	// Controlled error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return userDomain.User{}, projectErrors.ErrUserNotFound
		}

		// Unknown error
		return userDomain.User{}, result.Error
	}

	return user, nil
}
