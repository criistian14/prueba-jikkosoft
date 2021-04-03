package repository

import (
	"errors"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	"gorm.io/gorm"
)

func (repo inquiryRepository) FindInquiry(inquiryToSearch inquiryDomain.Inquiry, isDeleted bool) (inquiryDomain.Inquiry, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return inquiryToSearch, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	var inquiry inquiryDomain.Inquiry
	if isDeleted {
		db = db.Unscoped()
	}

	result := db.Where(&inquiryToSearch).First(&inquiry)
	// Controlled error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return inquiryDomain.Inquiry{}, projectErrors.ErrInquiryNotFound
		}

		// Unknown error
		return inquiryDomain.Inquiry{}, result.Error
	}

	return inquiry, nil
}
