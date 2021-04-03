package repository

import inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"

type InquiryRepository interface {
	FindInquiry(inquiry inquiryDomain.Inquiry, isDeleted bool) (inquiryDomain.Inquiry, error)
	SaveInquiry(inquiry inquiryDomain.Inquiry) (inquiryDomain.Inquiry, error)
}

type inquiryRepository struct {
}

// * Create and return new instance
func NewInquiryRepository() *inquiryRepository {
	return &inquiryRepository{}
}
