package errors

import "errors"

var (
	ErrPublicServiceNotFound    = errors.New("the public service was not found")
	ErrPublicServiceDuplicateID = errors.New("there is a role with the same registered id")
)
