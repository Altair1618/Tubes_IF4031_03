package ticketService

import (
	"fmt"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTicketByIdService(id uuid.UUID) utils.ResponseBody {
	db, _ := configs.GetGormClient()

	var ticket models.Ticket
	result := db.First(&ticket, "id = ?", id)

	if result.RowsAffected == 0 {
		return utils.ResponseBody{
			Code:    404,
			Message: "Ticket Not Found",
			Data:    nil,
		}
	}

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	} else {
		return utils.ResponseBody{
			Code:    200,
			Message: "Ticket Data Fetched Successfully",
			Data:    fiber.Map{"ticket": ticket},
		}
	}
}