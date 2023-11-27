package ticketService

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

func UpdateTicketService(id uuid.UUID, payload commonStructs.UpdateTicketServicePayload) utils.ResponseBody {
	validator := utils.CustomValidator{
		Validator: validator.New(),
	}

	if err := validator.Validate(payload); err != nil {
		return utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: utils.GetValidationErrorMessages(err)[0].Message,
		}
	}

	ticket := models.Ticket{
		Id:        id,
		Price:    payload.Price,
		EventId:  payload.EventId,
		SeatId:  payload.SeatId,
		Status:  payload.Status,
	}

	db, _ := configs.GetGormClient()

	result := db.Model(&ticket).Updates(models.Ticket{
		Price:    payload.Price,
		EventId:  payload.EventId,
		SeatId:  payload.SeatId,
		Status:  payload.Status,
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
				Message: "Ticket Not Found",
				Data:    nil,
			}
		}
		
		return utils.ResponseBody{
			Code:    200,
			Message: "Ticket Data Updated Successfully",
			Data:    nil,
		}
	}
}
