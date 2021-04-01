package main

import "github.com/criistian14/prueba-jikkosoft/src/config"

func main() {
	server := config.NewServer()

	server.InitVars()
	server.LoadConfigurations()
	server.InitModules()
	server.RunServer()
}
