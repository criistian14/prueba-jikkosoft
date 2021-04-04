package seeders

import (
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	inquiryRepository "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/repository"
	inquiryUseCase "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/usecase"
	"time"
)

func InquirySeeder() error {
	inquiryRepo := inquiryRepository.NewInquiryRepository()
	inquiryUsecase := inquiryUseCase.NewInquiryUsecase(inquiryRepo)

	// Inquiry 1
	title := "No deja pagar"
	message := "Message test"
	state := inquiryDomain.InProgressInquiry
	category := inquiryDomain.ClaimInquiry
	var userID uint = 1

	_, err := inquiryUsecase.SaveInquiry(inquiryDomain.Inquiry{
		Title:     &title,
		Message:   &message,
		State:     &state,
		Category:  &category,
		UserID:    &userID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	// Inquiry 2
	title = "Enviar factura a correo"
	message = "Message"
	state = inquiryDomain.ClosedInquiry
	category = inquiryDomain.PetitionInquiry
	userID = 1

	_, err = inquiryUsecase.SaveInquiry(inquiryDomain.Inquiry{
		Title:     &title,
		Message:   &message,
		State:     &state,
		Category:  &category,
		UserID:    &userID,
		CreatedAt: time.Now().AddDate(0, 0, -24),
	})
	if err != nil {
		return err
	}

	return nil
}
