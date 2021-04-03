package errors

import "errors"

var (
	ErrPublicServiceNotFound    = errors.New("the public service was not found")
	ErrPublicServiceDuplicateID = errors.New("there is a public service with the same registered id")
	ErrInvoiceNotFound          = errors.New("the invoice was not found")
	ErrInvoiceDuplicateID       = errors.New("there is an invoice with the same registered id")
)