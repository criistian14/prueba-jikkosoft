package seeders

import (
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
	userRepository "github.com/criistian14/prueba-jikkosoft/src/modules/users/repository"
	userUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/users/usecase"
	"time"
)

func UserSeeder() error {
	userRepo := userRepository.NewUserRepository()
	userUsecase := userUseCase.NewUserUsecase(userRepo)

	// User 1
	firstName := "Christian"
	lastName := "B"
	email := "cristian@hotmail.com"
	address := "Carrera"
	phone := "3191234567"
	var cityID uint = 1

	_, err := userUsecase.SaveUser(userDomain.User{
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     &email,
		Address:   &address,
		Phone:     &phone,
		CityID:    &cityID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	// User 2
	firstName = "Jack"
	lastName = "S"
	email = "jack@hotmail.com"
	address = "Avenida"
	phone = "3201234567"
	cityID = 2

	_, err = userUsecase.SaveUser(userDomain.User{
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     &email,
		Address:   &address,
		Phone:     &phone,
		CityID:    &cityID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
