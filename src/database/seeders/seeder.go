package seeders

import (
	"fmt"
	"github.com/criistian14/prueba-jikkosoft/src/database"
)

// Seeder database
func Seeder(forceSeeder bool) {
	dbClient, _ := database.NewDBClient()
	isSeeder := dbClient.GetStatusMigrate()

	if forceSeeder {
		isSeeder = true
	}

	if isSeeder {
		err := PublicServiceSeeder()
		if err != nil {
			fmt.Printf("Error Seeding: %v \n", err)
		}

		err = InvoiceSeeder()
		if err != nil {
			fmt.Printf("Error Seeding: %v \n", err)
		}

		err = InquirySeeder()
		if err != nil {
			fmt.Printf("Error Seeding: %v \n", err)
		}
	}
}
