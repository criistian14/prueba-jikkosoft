package repository

import (
	"errors"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	"gorm.io/gorm"
)

func (repo invoiceRepository) FindInvoice(invoiceToSearch invoiceDomain.Invoice, isDeleted bool) (invoiceDomain.Invoice, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return invoiceToSearch, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	var invoice invoiceDomain.Invoice
	if isDeleted {
		db = db.Unscoped()
	}

	result := db.Where(&invoiceToSearch).First(&invoice)
	// Controlled error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return invoiceDomain.Invoice{}, projectErrors.ErrInvoiceNotFound
		}

		// Unknown error
		return invoiceDomain.Invoice{}, result.Error
	}

	return invoice, nil
}
