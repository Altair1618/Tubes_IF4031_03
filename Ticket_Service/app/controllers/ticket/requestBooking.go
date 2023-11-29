package ticketController

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	bookingService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/booking"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestBookingController(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)
	
	if err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	requestPayload := new(commonStructs.RequestBookingControllerPayload)
	if err := c.BodyParser(requestPayload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	payload := new(commonStructs.RequestBookingServicePayload)
	payload.Token = c.Locals("token").(string)
	payload.BookingId = requestPayload.BookingId
	payload.TicketId = uuid
	payload.UserId = c.Locals("userInfo").(commonStructs.JWTPayload).UserId

	serviceResponse := bookingService.RequestBookingService(*payload)
	return utils.CreateResponseBody(c, serviceResponse)
}