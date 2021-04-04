package domain

import (
	"gorm.io/gorm"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type StateInquiry string

const (
	NewInquiry        StateInquiry = "new"
	InProgressInquiry StateInquiry = "in_progress"
	ClosedInquiry     StateInquiry = "closed"
)

type CategoryInquiry string

const (
	PetitionInquiry  CategoryInquiry = "petition"
	ComplaintInquiry CategoryInquiry = "complaint"
	ClaimInquiry     CategoryInquiry = "claim"
)

// * Inquiry
type Inquiry struct {
	ID       uint             `gorm:"primary_key" json:"id"`
	Title    *string          `gorm:"type:varchar(250); not null" json:"title"`
	Message  *string          `gorm:"type:varchar(250); not null" json:"message"`
	State    *StateInquiry    `gorm:"type:enum('new', 'in_progress', 'closed'); not null" json:"state"`
	Category *CategoryInquiry `gorm:"type:enum('petition', 'complaint', 'claim'); not null" json:"category"`
	UserID   *uint            `gorm:"not null" json:"user_id"`

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// * Set name to data base
func (inquiry Inquiry) TableName() string {
	return "inquiries"
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	Title    bool
	Message  bool
	State    bool
	Category bool
}

// * Validate data of domain
func (inquiry Inquiry) Validate(validate FieldsToValidate) error {
	var err error

	if validate.Title {
		err = validation.ValidateStruct(&inquiry,
			validation.Field(
				&inquiry.Title,
				validation.Required,
				validation.Length(3, 250),
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Message {
		err = validation.ValidateStruct(&inquiry,
			validation.Field(
				&inquiry.Message,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.State {
		err = validation.ValidateStruct(&inquiry,
			validation.Field(
				&inquiry.State,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	if validate.Category {
		err = validation.ValidateStruct(&inquiry,
			validation.Field(
				&inquiry.Category,
				validation.Required,
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}
