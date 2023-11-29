package routes

import (
	ticketController "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/controllers/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ticketRouteV1(v1 fiber.Router) {
	ticket := v1.Group("/tickets")

	ticket.Get("/", ticketController.GetManyTicketsByIdsController)
	ticket.Get("/:id", ticketController.GetTicketByIdController)
	ticket.Post("/", ticketController.CreateTicketController)
	ticket.Put("/:id", ticketController.UpdateTicketController)
	ticket.Patch("/:id", middlewares.AuthMiddleware, ticketController.UpdateStatusController)
	ticket.Delete("/:id", ticketController.DeleteTicketController)
}
