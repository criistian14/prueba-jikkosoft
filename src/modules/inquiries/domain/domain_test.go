package domain_test

import (
	"github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInquiry_Validate(t *testing.T) {
	title := "Title"
	message := "Title"
	state := domain.InProgressInquiry
	category := domain.PetitionInquiry

	data := domain.Inquiry{
		ID:       1,
		Title:    &title,
		Message:  &message,
		State:    &state,
		Category: &category,
	}

	err := data.Validate(domain.FieldsToValidate{
		Title:    true,
		Message:  true,
		State:    true,
		Category: true,
	})
	if err != nil {
		t.Fatalf("Error Validate Inquirys: %s", err)
	}

	// ! Errors
	t.Run("error-without-data", func(t *testing.T) {
		data := domain.Inquiry{}

		err := data.Validate(domain.FieldsToValidate{
			Title: true,
		})

		assert.Error(t, err)
	})
}
