package eventService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateEventService(id uuid.UUID, payload commonStructs.UpdateEventServicePayload) utils.ResponseBody {
	validator := utils.CustomValidator{
		Validator: validator.New(),
	}

	if err := validator.Validate(payload); err != nil {
		return utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: utils.GetValidationErrorMessages(err)[0].Message,
		}
	}

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

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Updating Data To Database",
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
			Message: "Event Data Updated Successfully",
			Data:    nil,
		}
	}
}
