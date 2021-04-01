package config

import (
	"fmt"
	"github.com/criistian14/prueba-jikkosoft/src/modules/numbers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

const defaultPort = "8080"

// * Server structure
type server struct {
	Host string
	Port string

	// App Fiber
	App *fiber.App
}

// ! Return server struct
func NewServer() *server {
	return &server{}
}

// * Return instance fiber
func (s *server) GetFiberApp() *fiber.App {
	return s.App
}

// * Initialize environment variables
func (s *server) InitVars() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error load .env:", err)
	}

	// Get Values of app
	s.Port = os.Getenv("APP_PORT")
	if s.Port == "" {
		s.Port = defaultPort
	}
}

// * Load configuration packages
func (s *server) LoadConfigurations() {
	// Init new instance fiber
	s.App = fiber.New()
}

// * Init modules
func (s *server) InitModules() {
	numbers.InitModule(s)
}

// * Run server
func (s *server) RunServer() {
	fmt.Println(s.App.Listen(":" + s.Port))
}
