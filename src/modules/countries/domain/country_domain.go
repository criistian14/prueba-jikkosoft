package domain

import (
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	"gorm.io/gorm"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// * Country
type Country struct {
	ID   uint    `gorm:"primary_key" json:"id"`
	Name *string `gorm:"type:varchar(250); not null" json:"name"`

	Cities []cityDomain.City `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

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
