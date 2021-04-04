package usecase

import (
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	inquiryRepository "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/repository"
)

type InquiryUsecase interface {
	GetInquiry(inquiry inquiryDomain.Inquiry, isDeleted bool) (inquiryDomain.Inquiry, error)
	SaveInquiry(inquiry inquiryDomain.Inquiry) (inquiryDomain.Inquiry, error)
}

type inquiryUsecase struct {
	repo inquiryRepository.InquiryRepository
}

// * Create and return new instance
func NewInquiryUsecase(repo inquiryRepository.InquiryRepository) *inquiryUsecase {
	return &inquiryUsecase{
		repo: repo,
	}
}
