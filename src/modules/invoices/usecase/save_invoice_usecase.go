package usecase

import (
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
)

func (usecase invoiceUsecase) SaveInvoice(invoice invoiceDomain.Invoice) (invoiceDomain.Invoice, error) {
	return usecase.repo.SaveInvoice(invoice)
}
