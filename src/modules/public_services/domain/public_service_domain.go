package domain

import (
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
	"time"

	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type TypePublicService string

const (
	PublicWaterService    TypePublicService = "water"
	PublicElectricService TypePublicService = "electric"
)

// * Public Service
type PublicService struct {
	ID      uint               `gorm:"primary_key" json:"id"`
	Company *string            `gorm:"type:varchar(250); not null" json:"company"`
	Type    *TypePublicService `gorm:"type:enum('water', 'electric'); not null" json:"type"`
	Email   *string            `gorm:"type:varchar(250); not null; unique" json:"email"`

	Invoices []invoiceDomain.Invoice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// * Set name to data base
func (ps PublicService) TableName() string {
	return "public_services"
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	Company bool
	Type    bool
	Email   bool
}

// * Validate data of domain
func (ps PublicService) Validate(validate FieldsToValidate) error {
	var err error

	if validate.Company {
		err = validation.ValidateStruct(&ps,
			validation.Field(
				&ps.Company, validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Email {
		err = validation.ValidateStruct(&ps,
			validation.Field(
				&ps.Email,
				validation.Required,
				is.Email,
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Type {
		err = validation.ValidateStruct(&ps,
			validation.Field(
				&ps.Type,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}
