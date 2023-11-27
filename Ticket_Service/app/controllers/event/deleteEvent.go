package eventController

import (
	eventService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/event"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteEventController(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	serviceResponse := eventService.DeleteEventService(uuid)
	return utils.CreateResponseBody(c, serviceResponse)
}