package ticketService

import (
	"fmt"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/google/uuid"
)

func DeleteTicketService(id uuid.UUID) utils.ResponseBody {
	ticket := models.Ticket{
		Id: id,
	}

	db, _ := configs.GetGormClient()

	result := db.Delete(&ticket)

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Deleting Data From Database",
			Data:    nil,
		}
	} else {
		if result.RowsAffected == 0 {
			return utils.ResponseBody{
				Code:    404,
				Message: "Ticket Not Found",
				Data:    nil,
			}
		}
		
		return utils.ResponseBody{
			Code:    200,
			Message: "Ticket Data Deleted Successfully",
			Data:    nil,
		}
	}
}