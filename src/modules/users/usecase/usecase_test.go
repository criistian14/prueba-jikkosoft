package usecase_test

import (
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
	"github.com/criistian14/prueba-jikkosoft/src/modules/users/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (repo *repositoryMock) SaveUser(_ userDomain.User) (userDomain.User, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(userDomain.User), args.Error(1)
}

func (repo *repositoryMock) FindUser(_ userDomain.User, _ bool) (userDomain.User, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(userDomain.User), args.Error(1)
}

func TestUserUsecase_SaveUser(t *testing.T) {
	mockRepo := new(repositoryMock)
	user := userDomain.User{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("SaveUser").Return(user, nil)
	testUsecase := usecase.NewUserUsecase(mockRepo)
	result, err := testUsecase.SaveUser(user)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}

func TestUserUsecase_GetUser(t *testing.T) {
	mockRepo := new(repositoryMock)
	user := userDomain.User{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("FindUser").Return(user, nil)
	testUsecase := usecase.NewUserUsecase(mockRepo)
	result, err := testUsecase.GetUser(user, false)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}
