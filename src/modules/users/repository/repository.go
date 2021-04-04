package repository

import userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"

type UserRepository interface {
	FindUser(user userDomain.User, isDeleted bool) (userDomain.User, error)
	SaveUser(user userDomain.User) (userDomain.User, error)
}

type userRepository struct {
}

// * Create and return new instance
func NewUserRepository() *userRepository {
	return &userRepository{}
}
