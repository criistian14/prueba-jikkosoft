package interfaces

import "github.com/gofiber/fiber/v2"

type Server interface {
	GetFiberApp() *fiber.App
}
