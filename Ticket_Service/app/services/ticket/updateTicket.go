package ticketService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	// "github.com/go-playground/validator/v10"
	// "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateTicketService(id uuid.UUID, payload commonStructs.UpdateTicketServicePayload) utils.ResponseBody {
	// validator := utils.CustomValidator{
	// 	Validator: validator.New(),
	// }

	// if err := validator.Validate(payload); err != nil {
	// 	return utils.ResponseBody{
	// 		Code:    fiber.StatusBadRequest,
	// 		Message: utils.GetValidationErrorMessages(err)[0].Message,
	// 	}
	// }

	// ticket := models.Ticket{
	// 	Id:      id,
	// }
	
	db, _ := configs.GetGormClient()
	
	var ticket models.Ticket
	result := db.First(&ticket, "id = ?", id)

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
			Message: "Ticket Not Found",
			Data:    nil,
		}
	}

	fmt.Println(payload)
	if payload.Price != 0 {
		ticket.Price = payload.Price
	}

	if payload.EventId != uuid.Nil {
		ticket.EventId = payload.EventId
	}

	if payload.SeatId != "" {
		ticket.SeatId = payload.SeatId
	}

	if payload.Status != "" {
		ticket.Status = payload.Status
	}

	result = db.Save(&ticket)

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
