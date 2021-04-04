package database

import (
	"fmt"
	cityDomain "github.com/criistian14/prueba-jikkosoft/src/modules/cities/domain"
	countryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/countries/domain"
	inquiryDomain "github.com/criistian14/prueba-jikkosoft/src/modules/inquiries/domain"
	invoiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/invoices/domain"
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
	userDomain "github.com/criistian14/prueba-jikkosoft/src/modules/users/domain"
)

// * Migrate tables
func Migrate(forceMigrate bool) {
	dbClient, _ := NewDBClient()
	isMigrate := dbClient.GetStatusMigrate()

	if forceMigrate {
		isMigrate = true
	}

	if isMigrate {
		db := dbClient.DB
		defer dbClient.CloseDataBase()

		// Drop tables
		err := db.Debug().Migrator().DropTable(
			&countryDomain.Country{},
			&cityDomain.City{},
			&userDomain.User{},
			&publicServiceDomain.PublicService{},
			&invoiceDomain.Invoice{},
			&inquiryDomain.Inquiry{},
		)
		if err != nil {
			fmt.Printf("Error Migrations: %v \n", err)
		}

		// Create tables
		err = db.Debug().Migrator().AutoMigrate(
			&countryDomain.Country{},
			&cityDomain.City{},
			&userDomain.User{},
			&publicServiceDomain.PublicService{},
			&invoiceDomain.Invoice{},
			&inquiryDomain.Inquiry{},
		)
		if err != nil {
			fmt.Printf("Error Migrations: %v \n", err)
		}
	}
}
