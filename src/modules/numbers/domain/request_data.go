package domain

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"reflect"
)

// * Request data
type RequestData struct {
	Numbers []int
}

// * Fields to validate of the domain
type FieldsToValidate struct {
	Numbers bool
}

// * Validate data of domain
func (data RequestData) Validate(validate FieldsToValidate) error {
	var err error

	if validate.Numbers {
		err = validation.ValidateStruct(&data,
			validation.Field(
				&data.Numbers,
				validation.Required.Error("The 'numbers' parameter is required"),
				validation.By(checkIsList),
			),
		)
	}
	if err != nil {
		return err
	}

	return nil
}

// ! Validate is a list of numbers
func checkIsList(value interface{}) error {
	typeValue := reflect.TypeOf(value).Kind()

	if typeValue == reflect.Slice || typeValue == reflect.Array {
		return nil
	}

	return fmt.Errorf("must be a list of numbers")
}
