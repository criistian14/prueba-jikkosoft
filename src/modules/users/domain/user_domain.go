package domain

import (
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// * User
type User struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	FirstName *string `gorm:"type:varchar(250); not null" json:"first_name"`
	LastName  *string `gorm:"type:varchar(250); not null" json:"last_name"`
	Email     *string `gorm:"type:varchar(300); not null" json:"email"`
	Address   *string `gorm:"type:varchar(250); not null" json:"address"`
	Phone     *string `gorm:"type:varchar(11); not null" json:"phone"`
	CityID    *uint   `gorm:"not null" json:"city_id"`

	Invoices  []invoiceDomain.Invoice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"invoices,omitempty"`
	Inquiries []inquiryDomain.Inquiry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"inquiries,omitempty"`

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// * Set name to data base
func (user User) TableName() string {
	return "users"
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	FirstName bool
	LastName  bool
	Email     bool
	Address   bool
	Phone     bool
	CityID    bool
}

// * Validate data of domain
func (user User) Validate(validate FieldsToValidate) error {
	var err error

	if validate.FirstName {
		err = validation.ValidateStruct(&user,
			validation.Field(
				&user.FirstName,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.LastName {
		err = validation.ValidateStruct(&user,
			validation.Field(
				&user.LastName,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Email {
		err = validation.ValidateStruct(&user,
			validation.Field(
				&user.Email,
				validation.Required,
				validation.Length(3, 300),
				is.Email,
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Address {
		err = validation.ValidateStruct(&user,
			validation.Field(
				&user.Address,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Phone {
		err = validation.ValidateStruct(&user,
			validation.Field(
				&user.Phone,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.CityID {
		err = validation.ValidateStruct(&user,
			validation.Field(
				&user.CityID,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}
