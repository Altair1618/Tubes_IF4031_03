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
	ticket.Put("/:id/book", middlewares.AuthMiddleware, ticketController.RequestBookingController)
	ticket.Patch("/", middlewares.AuthMiddleware, ticketController.UpdateStatusController)
	ticket.Patch("/:id/status/cancel", ticketController.CancelTicketController)
	ticket.Delete("/:id", ticketController.DeleteTicketController)
}
