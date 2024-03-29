package eventController

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	eventService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/event"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateEventController(c *fiber.Ctx) error {
	payload := new(commonStructs.CreateEventServicePayload)

	if err := c.BodyParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	serviceResponse := eventService.CreateEventService(*payload)
	return utils.CreateResponseBody(c, serviceResponse)
}
