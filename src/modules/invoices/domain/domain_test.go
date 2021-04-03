package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInvoice_Validate(t *testing.T) {
	paymentDeadline := time.Now()
	totalAmount := 200000
	var invoiceId uint = 1

	data := domain.Invoice{
		ID:              1,
		PaymentDeadline: &paymentDeadline,
		TotalAmount:     &totalAmount,
		PublicServiceID: &invoiceId,
	}

	err := data.Validate(domain.FieldsToValidate{
		PaymentDeadline: true,
		TotalAmount:     true,
		PublicServiceID: true,
	})
	if err != nil {
		t.Fatalf("Error Validate Invoices: %s", err)
	}

	// ! Errors
	t.Run("error-payment-deadline-passed", func(t *testing.T) {
		paymentDeadline := time.Now().AddDate(0, 0, -1)
		data := domain.Invoice{
			PaymentDeadline: &paymentDeadline,
		}

		err := data.Validate(domain.FieldsToValidate{
			PaymentDeadline: true,
		})

		assert.Error(t, err)
	})

	t.Run("error-without-data", func(t *testing.T) {
		data := domain.Invoice{}

		err := data.Validate(domain.FieldsToValidate{
			TotalAmount: true,
		})

		assert.Error(t, err)
	})
}
