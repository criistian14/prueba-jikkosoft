package usecase

import (
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
)

func (usecase invoiceUsecase) GetInvoice(invoice invoiceDomain.Invoice, isDeleted bool) (invoiceDomain.Invoice, error) {
	return usecase.repo.FindInvoice(invoice, isDeleted)
}
