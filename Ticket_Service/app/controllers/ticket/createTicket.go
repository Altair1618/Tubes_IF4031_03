package ticketController

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	ticketService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateTicketController(c *fiber.Ctx) error {
	payload := new(commonStructs.CreateTicketServicePayload)

	if err := c.BodyParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	serviceResponse := ticketService.CreateTicketService(*payload)
	return utils.CreateResponseBody(c, serviceResponse)
}