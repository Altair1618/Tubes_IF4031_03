package routes

import (
	eventController "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/controllers/event"
	"github.com/gofiber/fiber/v2"
)

func eventRouteV1(v1 fiber.Router) {
	event := v1.Group("/event")

	event.Get("/", eventController.GetEventsController)
	event.Post("/", eventController.CreateEventController)
}