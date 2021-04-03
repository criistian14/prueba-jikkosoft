package usecase

import (
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
)

func (usecase inquiryUsecase) GetInquiry(inquiry inquiryDomain.Inquiry, isDeleted bool) (inquiryDomain.Inquiry, error) {
	return usecase.repo.FindInquiry(inquiry, isDeleted)
}
