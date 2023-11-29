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

func GetEventByIdService(id uuid.UUID) utils.ResponseBody {
	db, _ := configs.GetGormClient()
	eventResponse := new(commonStructs.EventDetailResponse)

	var event models.Event
	result := db.First(&event, "id = ?", id)

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	}

	if result.RowsAffected == 0 {
		return utils.ResponseBody{
			Code:    404,
			Message: "Event Not Found",
			Data:    nil,
		}
	}

	var availableSeats int64
	db.Model(&models.Ticket{}).Where("event_id = ? AND status = ?", id, commonStructs.Open).Count(&availableSeats)
	
	eventResponse.Id = event.Id
	eventResponse.EventName = event.EventName
	eventResponse.EventTime = event.EventTime
	eventResponse.Location = event.Location
	eventResponse.CreatedAt = event.CreatedAt
	eventResponse.UpdatedAt = event.UpdatedAt
	eventResponse.AvailableSeats = int(availableSeats)

	return utils.ResponseBody{
		Code:    200,
		Message: "Event Data Fetched Successfully",
		Data:    fiber.Map{"event": eventResponse},
	}
}