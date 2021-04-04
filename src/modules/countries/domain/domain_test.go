package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountry_Validate(t *testing.T) {
	name := "Colombia"

	data := domain.Country{
		ID:   1,
		Name: &name,
	}

	err := data.Validate(domain.FieldsToValidate{
		Name: true,
	})
	if err != nil {
		t.Fatalf("Error Validate Countrys: %s", err)
	}

	// ! Errors
	t.Run("error-without-data", func(t *testing.T) {
		data := domain.Country{}

		err := data.Validate(domain.FieldsToValidate{
			Name: true,
		})

		assert.Error(t, err)
	})
}
