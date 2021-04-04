package usecase_test

import (
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	"github.com/criistian14/prueba-jikkosoft/src/modules/cities/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (repo *repositoryMock) SaveCity(_ cityDomain.City) (cityDomain.City, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(cityDomain.City), args.Error(1)
}

func (repo *repositoryMock) FindCity(_ cityDomain.City, _ bool) (cityDomain.City, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(cityDomain.City), args.Error(1)
}

func TestCityUsecase_SaveCity(t *testing.T) {
	mockRepo := new(repositoryMock)
	city := cityDomain.City{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("SaveCity").Return(city, nil)
	testUsecase := usecase.NewCityUsecase(mockRepo)
	result, err := testUsecase.SaveCity(city)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}

func TestCityUsecase_GetCity(t *testing.T) {
	mockRepo := new(repositoryMock)
	city := cityDomain.City{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("FindCity").Return(city, nil)
	testUsecase := usecase.NewCityUsecase(mockRepo)
	result, err := testUsecase.GetCity(city, false)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}
