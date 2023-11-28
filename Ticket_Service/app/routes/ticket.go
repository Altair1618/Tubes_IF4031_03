package routes

import (
	ticketController "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/controllers/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ticketRouteV1(v1 fiber.Router) {
	ticket := v1.Group("/ticket")

	ticket.Get("/:id", ticketController.GetTicketByIdController)
	ticket.Get("/ids", ticketController.GetManyTicketsByIdsController)
	ticket.Post("/", ticketController.CreateTicketController)
	ticket.Put("/:id", ticketController.UpdateTicketController)
	ticket.Patch("/", middlewares.AuthMiddleware, ticketController.UpdateStatusController)
	ticket.Delete("/:id", ticketController.DeleteTicketController)
}
