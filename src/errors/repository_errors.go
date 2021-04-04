package errors

import "errors"

var (
	ErrPublicServiceNotFound    = errors.New("the public service was not found")
	ErrPublicServiceDuplicateID = errors.New("there is a public service with the same registered id")
	ErrInvoiceNotFound          = errors.New("the invoice was not found")
	ErrInvoiceDuplicateID       = errors.New("there is an invoice with the same registered id")
	ErrInquiryNotFound          = errors.New("the inquiry was not found")
	ErrInquiryDuplicateID       = errors.New("there is an inquiry with the same registered id")
	ErrCountryNotFound          = errors.New("the country was not found")
	ErrCountryDuplicateID       = errors.New("there is an country with the same registered id")
	ErrCityNotFound             = errors.New("the city was not found")
	ErrCityDuplicateID          = errors.New("there is an city with the same registered id")
	ErrUserNotFound             = errors.New("the user was not found")
	ErrUserDuplicateID          = errors.New("there is an user with the same registered id")
)
