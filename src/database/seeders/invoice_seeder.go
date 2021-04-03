package seeders

import (
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	invoiceRepository "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/repository"
	invoiceUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/usecase"
	"time"
)

func InvoiceSeeder() error {
	invoiceRepo := invoiceRepository.NewInvoiceRepository()
	invoiceUsecase := invoiceUseCase.NewInvoiceUsecase(invoiceRepo)

	// Invoice 1
	paymentDeadline := time.Now().AddDate(0, 1, 0)
	totalAmount := 200000
	paid := false
	var publicServiceID uint = 1

	_, err := invoiceUsecase.SaveInvoice(invoiceDomain.Invoice{
		PaymentDeadline: &paymentDeadline,
		TotalAmount:     &totalAmount,
		Paid:            &paid,
		PublicServiceID: &publicServiceID,
		CreatedAt:       time.Now(),
	})
	if err != nil {
		return err
	}

	// Invoice 2
	paymentDeadline = time.Now().AddDate(0, 0, 1)
	totalAmount = 167000
	paid = true
	publicServiceID = 2

	_, err = invoiceUsecase.SaveInvoice(invoiceDomain.Invoice{
		PaymentDeadline: &paymentDeadline,
		TotalAmount:     &totalAmount,
		Paid:            &paid,
		PublicServiceID: &publicServiceID,
		CreatedAt:       time.Now().AddDate(0, 0, -24),
	})
	if err != nil {
		return err
	}

	return nil
}
