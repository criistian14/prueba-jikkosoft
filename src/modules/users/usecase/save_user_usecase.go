package usecase

import (
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
)

func (usecase userUsecase) SaveUser(user userDomain.User) (userDomain.User, error) {
	return usecase.repo.SaveUser(user)
}
