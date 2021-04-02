package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicService_Validate(t *testing.T) {
	company := "Company Test"
	typePublicService := domain.PublicWaterService
	email := "test@test.com"

	data := domain.PublicService{
		ID:      1,
		Company: &company,
		Type:    &typePublicService,
		Email:   &email,
	}

	err := data.Validate(domain.FieldsToValidate{
		Company: true,
		Type:    true,
		Email:   true,
	})
	if err != nil {
		t.Fatalf("Error Validate PublicServices: %s", err)
	}

	// ! Errors
	t.Run("error-without-data", func(t *testing.T) {
		data := domain.PublicService{}

		err := data.Validate(domain.FieldsToValidate{
			Company: true,
		})

		assert.Error(t, err)
	})
}
