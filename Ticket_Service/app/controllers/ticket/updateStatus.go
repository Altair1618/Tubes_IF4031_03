package ticketController

import (
	ticketService "github.com/Altair1618/IF4031_03_Ticket/app/services/ticket"
	"github.com/Altair1618/IF4031_03_Ticket/app/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdateStatusController(c *fiber.Ctx) error {

	serviceResponse := ticketService.UpdateStatusService()
	return utils.CreateResponseBody(c, serviceResponse)
}
