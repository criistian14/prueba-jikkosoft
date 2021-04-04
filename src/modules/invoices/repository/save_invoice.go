package repository

import (
	projectConditionals "github.com/criistian14/prueba-jikkosoft/src/conditionals"
	"github.com/criistian14/prueba-jikkosoft/src/database"
	projectErrors "github.com/criistian14/prueba-jikkosoft/src/errors"
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
)

func (repo invoiceRepository) SaveInvoice(invoice invoiceDomain.Invoice) (invoiceDomain.Invoice, error) {
	dbClient, err := database.NewDBClient()
	if err != nil {
		return invoice, err
	}
	db := dbClient.DB
	defer dbClient.CloseDataBase()

	result := db.Create(&invoice)
	// Controlled error
	if result.Error != nil {
		// Duplicate ID
		if projectConditionals.IDIsDuplicated(result.Error) {
			return invoice, projectErrors.ErrInvoiceDuplicateID
		}

		// Unknown error
		return invoice, result.Error
	}

	return invoice, nil
}
