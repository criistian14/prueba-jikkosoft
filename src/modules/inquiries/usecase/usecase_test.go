package usecase_test

import (
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	"github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (repo *repositoryMock) SaveInquiry(_ inquiryDomain.Inquiry) (inquiryDomain.Inquiry, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(inquiryDomain.Inquiry), args.Error(1)
}

func (repo *repositoryMock) FindInquiry(_ inquiryDomain.Inquiry, _ bool) (inquiryDomain.Inquiry, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(inquiryDomain.Inquiry), args.Error(1)
}

func TestInquiryUsecase_SaveInquiry(t *testing.T) {
	mockRepo := new(repositoryMock)
	inquiry := inquiryDomain.Inquiry{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("SaveInquiry").Return(inquiry, nil)
	testUsecase := usecase.NewInquiryUsecase(mockRepo)
	result, err := testUsecase.SaveInquiry(inquiry)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}

func TestInquiryUsecase_GetInquiry(t *testing.T) {
	mockRepo := new(repositoryMock)
	inquiry := inquiryDomain.Inquiry{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("FindInquiry").Return(inquiry, nil)
	testUsecase := usecase.NewInquiryUsecase(mockRepo)
	result, err := testUsecase.GetInquiry(inquiry, false)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}
