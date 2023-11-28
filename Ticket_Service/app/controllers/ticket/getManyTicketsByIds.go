package ticketController

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	ticketService "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/services/ticket"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetManyTicketsByIdsController(c *fiber.Ctx) error {

	payload := new (commonStructs.GetManyTicketsByIdsPayload)

	if err := c.QueryParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid list of ids",
		})
	}

	// Get the "ids" parameter from the URL
	ids := payload.Ids

	// Convert the strings to UUIDs
	uuids := make([]uuid.UUID, len(ids))
	for i, id := range ids {
		parsedUUID, err := uuid.Parse(id)
		if err != nil {
			return utils.CreateResponseBody(c, utils.ResponseBody{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid id in ids parameter",
			})
		}
		uuids[i] = parsedUUID
	}

	// Call the service with the array of UUIDs
	serviceResponse := ticketService.GetManyTicketsByIdsService(uuids)
	return utils.CreateResponseBody(c, serviceResponse)
}