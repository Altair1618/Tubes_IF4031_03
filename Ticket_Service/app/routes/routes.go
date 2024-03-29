package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Static("/", "./public")

	api := app.Group("/api")

	// version 1 group
	v1 := api.Group("/v1")
	eventRouteV1(v1)
	ticketRouteV1(v1)
}
