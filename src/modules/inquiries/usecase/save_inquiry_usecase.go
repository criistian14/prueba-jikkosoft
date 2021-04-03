package usecase

import (
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
)

func (usecase inquiryUsecase) SaveInquiry(inquiry inquiryDomain.Inquiry) (inquiryDomain.Inquiry, error) {
	return usecase.repo.SaveInquiry(inquiry)
}
