package ticketController

import (
	commonStructs "github.com/Altair1618/IF4031_03_Ticket/app/common/structs"
	ticketService "github.com/Altair1618/IF4031_03_Ticket/app/services/ticket"
	"github.com/Altair1618/IF4031_03_Ticket/app/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdateStatusController(c *fiber.Ctx) error {
	payload := new(commonStructs.UpdateStatusServicePayload)

	if err := c.BodyParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	serviceResponse := ticketService.UpdateStatusService(*payload)
	return utils.CreateResponseBody(c, serviceResponse)
}
