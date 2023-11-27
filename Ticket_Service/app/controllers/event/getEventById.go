package eventController

import (
	eventService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/event"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetEventByIdController(c *fiber.Ctx) error {
	id := c.Params("id")

	serviceResponse := eventService.GetEventByIdService(id)
	return utils.CreateResponseBody(c, serviceResponse)
}