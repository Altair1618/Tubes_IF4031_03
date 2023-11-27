package eventService

import (
	"fmt"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/google/uuid"
)

func DeleteEventService(id uuid.UUID) utils.ResponseBody {
	event := models.Event{
		Id: id,
	}

	db, _ := configs.GetGormClient()

	result := db.Delete(&event)

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
				Message: "Event Not Found",
				Data:    nil,
			}
		}
	
		return utils.ResponseBody{
			Code:    200,
			Message: "Event Data Deleted Successfully",
			Data:    nil,
		}
	}
}