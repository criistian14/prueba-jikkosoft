package database

import (
	"fmt"
	publicServiceDomain "github.com/criistian14/prueba-jikkosoft/src/modules/public_services/domain"
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
			&publicServiceDomain.PublicService{},
		)
		if err != nil {
			fmt.Printf("Error Migrations: %v \n", err)
		}

		// Create tables
		err = db.Debug().Migrator().AutoMigrate(
			&publicServiceDomain.PublicService{},
		)
		if err != nil {
			fmt.Printf("Error Migrations: %v \n", err)
		}
	}
}
