package repository

import (
	projectConditionals "github.com/criistian14/prueba-jikkosoft/src/conditionals"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
)

func (repo userRepository) SaveUser(user userDomain.User) (userDomain.User, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return user, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	result := db.Create(&user)
	// Controlled error
	if result.Error != nil {
		// Duplicate ID
		if projectConditionals.IDIsDuplicated(result.Error) {
			return user, projectErrors.ErrUserDuplicateID
		}

		// Unknown error
		return user, result.Error
	}

	return user, nil
}
