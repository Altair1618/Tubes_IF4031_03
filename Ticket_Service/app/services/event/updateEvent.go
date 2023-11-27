package eventService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateEventService(id uuid.UUID, payload commonStructs.UpdateEventServicePayload) utils.ResponseBody {
	event := models.Event{
		Id:        id,
		EventName: payload.EventName,
		EventTime: payload.EventTime,
		Location:  payload.Location,
	}

	db, _ := configs.GetGormClient()

	result := db.Model(&event).Updates(models.Event{
		EventName: payload.EventName,
		EventTime: payload.EventTime,
		Location:  payload.Location,
	})

	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Updating Data To Database",
			Data:    nil,
		}
	} else {
		// fmt.Println(event)

		return utils.ResponseBody{
			Code:    200,
			Message: "Event Data Updated Successfully",
			Data:    fiber.Map{"event": event},
		}
	}
}
