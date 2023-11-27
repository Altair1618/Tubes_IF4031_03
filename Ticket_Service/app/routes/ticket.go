package routes

import (
	ticketController "github.com/Altair1618/IF4031_03_Ticket/app/controllers/ticket"
	"github.com/gofiber/fiber/v2"
)

func ticketRouteV1(v1 fiber.Router) {
	ticket := v1.Group("/ticket")

	ticket.Get("/", ticketController.GetTicketController)
}
