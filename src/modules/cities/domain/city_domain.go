package domain

import (
	"gorm.io/gorm"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// * City
type City struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	Name      *string `gorm:"type:varchar(250); not null" json:"name"`
	CountryID *uint   `gorm:"not null" json:"country_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// * Set name to data base
func (city City) TableName() string {
	return "cities"
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	Name      bool
	CountryID bool
}

// * Validate data of domain
func (city City) Validate(validate FieldsToValidate) error {
	var err error

	if validate.Name {
		err = validation.ValidateStruct(&city,
			validation.Field(
				&city.Name,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.CountryID {
		err = validation.ValidateStruct(&city,
			validation.Field(
				&city.CountryID,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}
