package usecase

import (
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
)

func (usecase userUsecase) GetUser(user userDomain.User, isDeleted bool) (userDomain.User, error) {
	return usecase.repo.FindUser(user, isDeleted)
}
