package handler

import (
	invoiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/repository"
	invoiceUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/usecase"
)

type InvoiceHandler interface {
}

type invoiceHandler struct {
	invoiceUsecase invoiceUseCase.InvoiceUsecase
}

// * Create and return new instance
func NewInvoiceHandler() *invoiceHandler {
	invoiceRepo := invoiceRepository.NewInvoiceRepository()
	invoiceUsecase := invoiceUseCase.NewInvoiceUsecase(invoiceRepo)

	return &invoiceHandler{
		invoiceUsecase: invoiceUsecase,
	}
}
