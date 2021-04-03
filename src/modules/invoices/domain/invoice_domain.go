package domain

import (
	"gorm.io/gorm"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// * Invoice
type Invoice struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	PaymentDeadline *time.Time `gorm:"not null" json:"payment_deadline"`
	TotalAmount     *int       `gorm:"not null" json:"total_amount"`
	Paid            *bool      `gorm:"default:false; not null" json:"paid"`
	PublicServiceID *uint      `gorm:"not null" json:"public_service_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// * Set name to data base
func (invoice Invoice) TableName() string {
	return "invoices"
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	PaymentDeadline bool
	TotalAmount     bool
	PublicServiceID bool
}

// * Validate data of domain
func (invoice Invoice) Validate(validate FieldsToValidate) error {
	var err error

	if validate.PaymentDeadline {
		err = validation.ValidateStruct(&invoice,
			validation.Field(
				&invoice.PaymentDeadline,
				validation.Required,
				validation.Min(time.Now()),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.TotalAmount {
		err = validation.ValidateStruct(&invoice,
			validation.Field(
				&invoice.TotalAmount,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.PublicServiceID {
		err = validation.ValidateStruct(&invoice,
			validation.Field(
				&invoice.PublicServiceID,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}
