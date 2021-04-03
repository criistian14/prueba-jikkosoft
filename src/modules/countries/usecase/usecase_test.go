package usecase_test

import (
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
	"github.com/criistian14/prueba-jikkosoft/src/modules/countries/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (repo *repositoryMock) SaveCountry(_ countryDomain.Country) (countryDomain.Country, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(countryDomain.Country), args.Error(1)
}

func (repo *repositoryMock) FindCountry(_ countryDomain.Country, _ bool) (countryDomain.Country, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(countryDomain.Country), args.Error(1)
}

func TestCountryUsecase_SaveCountry(t *testing.T) {
	mockRepo := new(repositoryMock)
	country := countryDomain.Country{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("SaveCountry").Return(country, nil)
	testUsecase := usecase.NewCountryUsecase(mockRepo)
	result, err := testUsecase.SaveCountry(country)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}

func TestCountryUsecase_GetCountry(t *testing.T) {
	mockRepo := new(repositoryMock)
	country := countryDomain.Country{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("FindCountry").Return(country, nil)
	testUsecase := usecase.NewCountryUsecase(mockRepo)
	result, err := testUsecase.GetCountry(country, false)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}
