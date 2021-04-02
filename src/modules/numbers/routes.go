package numbers

import (
	"github.com/criistian14/prueba-jikkosoft/src/interfaces"
	numberHandler "github.com/criistian14/prueba-jikkosoft/src/modules/numbers/handler"
	"github.com/gofiber/fiber/v2"
)

func routes(server interfaces.Server) {
	fiberApp := server.GetFiberApp()
	api := fiberApp.Group("/numbers")
	handler := numberHandler.NewNumberHandler()

	api.Post("/sort", func(ctx *fiber.Ctx) error {
		return handler.SortArrayHandler(ctx)
	})
}
