package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api")

	// version 1 group
	v1 := api.Group("/v1")
	paymentRouteV1(v1)
}
