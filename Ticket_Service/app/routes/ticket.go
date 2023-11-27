package routes

import (
	ticketController "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/controllers/ticket"
	"github.com/gofiber/fiber/v2"
)

func ticketRouteV1(v1 fiber.Router) {
	ticket := v1.Group("/ticket")

	ticket.Get("/", ticketController.GetTicketController)
	ticket.Post("/", ticketController.CreateTicketController)
	ticket.Patch("/", ticketController.UpdateStatusController)
}
