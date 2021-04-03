package usecase_test

import (
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	"github.com/criistian14/prueba-jikkosoft/src/modules/invoices/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (repo *repositoryMock) SaveInvoice(_ invoiceDomain.Invoice) (invoiceDomain.Invoice, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(invoiceDomain.Invoice), args.Error(1)
}

func (repo *repositoryMock) FindInvoice(_ invoiceDomain.Invoice, _ bool) (invoiceDomain.Invoice, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(invoiceDomain.Invoice), args.Error(1)
}

func TestInvoiceUsecase_SaveInvoice(t *testing.T) {
	mockRepo := new(repositoryMock)
	invoice := invoiceDomain.Invoice{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("SaveInvoice").Return(invoice, nil)
	testUsecase := usecase.NewInvoiceUsecase(mockRepo)
	result, err := testUsecase.SaveInvoice(invoice)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}

func TestInvoiceUsecase_GetInvoice(t *testing.T) {
	mockRepo := new(repositoryMock)
	invoice := invoiceDomain.Invoice{
		ID: 1,
	}

	// Setup expectations
	mockRepo.On("FindInvoice").Return(invoice, nil)
	testUsecase := usecase.NewInvoiceUsecase(mockRepo)
	result, err := testUsecase.GetInvoice(invoice, false)

	// Not error
	assert.Nil(t, err)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, uint(1), result.ID)
}
