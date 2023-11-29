package ticketController

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	ticketService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UpdateStatusController(c *fiber.Ctx) error {

	payload := new(commonStructs.UpdateTicketStatusRequest)

	if err := c.BodyParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	fmt.Println(payload)

	validator := utils.CustomValidator{
		Validator: validator.New(),
	}
	if err := validator.Validate(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: utils.GetValidationErrorMessages(err)[0].Message,
		})
	}

	serviceResponse := ticketService.UpdateStatusService(commonStructs.UpdateTicketStatusServicePayload{
		InvoiceId: payload.InvoiceId,
		Status:    payload.Status,
		UserId:    c.Locals("userInfo").(commonStructs.JWTPayload).UserId,
		Message:   payload.Message,
	})

	return utils.CreateResponseBody(c, serviceResponse)
}
