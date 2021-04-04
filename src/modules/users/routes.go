package users

import (
	"github.com/criistian14/prueba-jikkosoft/src/interfaces"
	userHandler "github.com/criistian14/prueba-jikkosoft/src/modules/users/handler"
	"github.com/gofiber/fiber/v2"
)

func routes(server interfaces.Server) {
	fiberApp := server.GetFiberApp()
	api := fiberApp.Group("/users")
	handler := userHandler.NewUserHandler()

	api.Get("/:id", func(ctx *fiber.Ctx) error {
		return handler.FindUserHandler(ctx)
	})
}
