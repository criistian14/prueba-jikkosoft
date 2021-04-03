package handler

import (
	inquiryRepository "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/repository"
	inquiryUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/usecase"
)

type InquiryHandler interface {
}

type inquiryHandler struct {
	inquiryUsecase inquiryUseCase.InquiryUsecase
}

// * Create and return new instance
func NewInquiryHandler() *inquiryHandler {
	inquiryRepo := inquiryRepository.NewInquiryRepository()
	inquiryUsecase := inquiryUseCase.NewInquiryUsecase(inquiryRepo)

	return &inquiryHandler{
		inquiryUsecase: inquiryUsecase,
	}
}
