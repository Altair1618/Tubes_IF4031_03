package eventController

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	eventService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/event"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetEventsController(c *fiber.Ctx) error {
	payload := new(commonStructs.GetEventsServicePayload)

	if err := c.QueryParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code: fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	serviceResponse := eventService.GetEventsService(*payload)
	return utils.CreateResponseBody(c, serviceResponse)
}
