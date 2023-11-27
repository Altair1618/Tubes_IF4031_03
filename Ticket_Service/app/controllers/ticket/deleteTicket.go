package ticketController

import (
	ticketService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteTicketController(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	serviceResponse := ticketService.DeleteTicketService(uuid)
	return utils.CreateResponseBody(c, serviceResponse)
}