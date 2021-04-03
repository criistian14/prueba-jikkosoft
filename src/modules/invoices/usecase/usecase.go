package usecase

import (
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	invoiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/repository"
)

type InvoiceUsecase interface {
	GetInvoice(invoice invoiceDomain.Invoice, isDeleted bool) (invoiceDomain.Invoice, error)
	SaveInvoice(invoice invoiceDomain.Invoice) (invoiceDomain.Invoice, error)
}

type invoiceUsecase struct {
	repo invoiceRepository.InvoiceRepository
}

// * Create and return new instance
func NewInvoiceUsecase(repo invoiceRepository.InvoiceRepository) *invoiceUsecase {
	return &invoiceUsecase{
		repo: repo,
	}
}
