package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCity_Validate(t *testing.T) {
	name := "Cali"

	data := domain.City{
		ID:   1,
		Name: &name,
	}

	err := data.Validate(domain.FieldsToValidate{
		Name: true,
	})
	if err != nil {
		t.Fatalf("Error Validate Citys: %s", err)
	}

	// ! Errors
	t.Run("error-without-data", func(t *testing.T) {
		data := domain.City{}

		err := data.Validate(domain.FieldsToValidate{
			Name: true,
		})

		assert.Error(t, err)
	})
}
