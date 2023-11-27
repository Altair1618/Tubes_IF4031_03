package ticketController

import (
	ticketService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdateStatusController(c *fiber.Ctx) error {

	serviceResponse := ticketService.UpdateStatusService()
	return utils.CreateResponseBody(c, serviceResponse)
}
