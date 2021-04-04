package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	name := "Jack"

	data := domain.User{
		ID:        1,
		FirstName: &name,
	}

	err := data.Validate(domain.FieldsToValidate{
		FirstName: true,
	})
	if err != nil {
		t.Fatalf("Error Validate Users: %s", err)
	}

	// ! Errors
	t.Run("error-without-data", func(t *testing.T) {
		data := domain.User{}

		err := data.Validate(domain.FieldsToValidate{
			FirstName: true,
		})

		assert.Error(t, err)
	})
}
