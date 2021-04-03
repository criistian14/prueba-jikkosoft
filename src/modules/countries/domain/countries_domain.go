package domain

import (
	"gorm.io/gorm"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// * Country
type Country struct {
	ID   uint    `gorm:"primary_key" json:"id"`
	Name *string `gorm:"type:varchar(250); not null" json:"name"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// * Set name to data base
func (country Country) TableName() string {
	return "countries"
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	Name bool
}

// * Validate data of domain
func (country Country) Validate(validate FieldsToValidate) error {
	var err error

	if validate.Name {
		err = validation.ValidateStruct(&country,
			validation.Field(
				&country.Name,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}
