package invoiceController

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	invoiceService "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/services/invoice"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateInvoiceController(c *fiber.Ctx) error {
	requestPayload := new(commonStructs.CreateInvoiceControllerPayload)

	if err := c.BodyParser(requestPayload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code: fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	payload := new(commonStructs.CreateInvoiceServicePayload)
	payload.TicketId = requestPayload.TicketId
	payload.UserId = c.Locals("userInfo").(commonStructs.JWTPayload).UserId

	serviceResponse := invoiceService.CreateInvoiceService(*payload)
	return utils.CreateResponseBody(c, serviceResponse)
}