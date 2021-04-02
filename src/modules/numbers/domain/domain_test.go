package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/numbers/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestData_Validate(t *testing.T) {
	numbers := []int{123, 4, 5}
	data := domain.RequestData{
		Numbers: numbers,
	}

	err := data.Validate(domain.FieldsToValidate{
		Numbers: true,
	})
	if err != nil {
		t.Fatalf("Error Validate Numbers: %s", err)
	}

	// ! Errors
	t.Run("error-without-data", func(t *testing.T) {
		data := domain.RequestData{}

		err := data.Validate(domain.FieldsToValidate{
			Numbers: true,
		})

		assert.Error(t, err)
	})
}
