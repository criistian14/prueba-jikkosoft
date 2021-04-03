package repository

import invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"

type InvoiceRepository interface {
	FindInvoice(invoice invoiceDomain.Invoice, isDeleted bool) (invoiceDomain.Invoice, error)
	SaveInvoice(invoice invoiceDomain.Invoice) (invoiceDomain.Invoice, error)
}

type invoiceRepository struct {
}

// * Create and return new instance
func NewInvoiceRepository() *invoiceRepository {
	return &invoiceRepository{}
}
