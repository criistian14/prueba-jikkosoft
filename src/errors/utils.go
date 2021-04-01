package errors

import "errors"

var (
	ErrNumbersEmpty = errors.New("list numbers are empty")
	ErrNumbersNull  = errors.New("list numbers is null")
)
