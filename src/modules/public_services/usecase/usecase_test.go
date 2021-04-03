package usecase_test

import (
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	"github.com/criistian14/prueba-jikkosoft/src/modules/public_services/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (repo *repositoryMock) SavePublicService(_ publicServiceDomain.PublicService) (publicServiceDomain.PublicService, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(publicServiceDomain.PublicService), args.Error(1)
}

func (repo *repositoryMock) FindPublicService(_ publicServiceDomain.PublicService, _ bool) (publicServiceDomain.PublicService, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(publicServiceDomain.PublicService), args.Error(1)
}

func TestPublicServiceUsecase_SavePublicService(t *testing.T) {
	mockRepo := new(repositoryMock)
	publicService := publicServiceDomain.PublicService{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("SavePublicService").Return(publicService, nil)
	testUsecase := usecase.NewPublicServiceUsecase(mockRepo)
	result, err := testUsecase.SavePublicService(publicService)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}

func TestPublicServiceUsecase_GetPublicService(t *testing.T) {
	mockRepo := new(repositoryMock)
	publicService := publicServiceDomain.PublicService{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("FindPublicService").Return(publicService, nil)
	testUsecase := usecase.NewPublicServiceUsecase(mockRepo)
	result, err := testUsecase.GetPublicService(publicService, false)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}
