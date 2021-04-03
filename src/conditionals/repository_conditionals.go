package conditionals

import "strings"

// * If the id is duplicated
func IDIsDuplicated(err error) bool {
	return strings.Contains(err.Error(), "PRIMARY") || strings.Contains(err.Error(), "pkey") || strings.Contains(err.Error(), "UNIQUE")
}

// * If the email is duplicated
func EmailIsDuplicated(err error) bool {
	return strings.Contains(err.Error(), "email")
}

// * If the company is duplicated
func CompanyIsDuplicated(err error) bool {
	return strings.Contains(err.Error(), "company")
}
