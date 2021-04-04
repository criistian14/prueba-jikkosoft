package main

import (
	"github.com/criistian14/prueba-jikkosoft/src/config"
	_ "github.com/joho/godotenv/autoload"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

func main() {
	server := config.NewServer()

	server.InitVars()
	server.LoadConfigurations()
	server.InitModules()
	server.RunServer()
}
