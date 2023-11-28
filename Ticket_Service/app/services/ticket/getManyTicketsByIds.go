package ticketService

import (
	"fmt"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetManyTicketsByIdsService(ids []uuid.UUID) utils.ResponseBody {
	db, _ := configs.GetGormClient()

	var tickets []models.Ticket
	result := db.Find(&tickets, "id IN ?", ids)

	if result.Error != nil {
		fmt.Println(result.Error)
		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	}

	// Organize the fetched tickets in a map with the IDs as keys
	ticketMap := make(map[uuid.UUID]models.Ticket)
	for _, ticket := range tickets {
		ticketMap[ticket.Id] = ticket
	}

	return utils.ResponseBody{
		Code:    200,
		Message: "Tickets Data Fetched Successfully",
		Data:    fiber.Map{"tickets": ticketMap},
	}
}