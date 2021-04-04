package usecase

import (
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
	userRepository "github.com/criistian14/prueba-jikkosoft/src/modules/users/repository"
)

type UserUsecase interface {
	GetUser(user userDomain.User, isDeleted bool) (userDomain.User, error)
	SaveUser(user userDomain.User) (userDomain.User, error)
}

type userUsecase struct {
	repo userRepository.UserRepository
}

// * Create and return new instance
func NewUserUsecase(repo userRepository.UserRepository) *userUsecase {
	return &userUsecase{
		repo: repo,
	}
}
