package bookingService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestBookingService(ticketId uuid.UUID) utils.ResponseBody {
	// Check if ticket exists
	db, _ := configs.GetGormClient()

	var ticket models.Ticket
	result := db.First(&ticket, "id = ?", ticketId)

	if result.Error != nil {
		fmt.Println(result.Error)
		
		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	}

	if result.RowsAffected == 0 {
		return utils.ResponseBody{
			Code:    fiber.StatusNotFound,
			Message: "Ticket Not Found",
			Data:    nil,
		}
	}

	// Simulate Failed External Call
	if utils.SimulateProbability(20) {
		// TODO: Generate PDF

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed To Book Ticket",
			Data:    fiber.Map{"status": "failed"},
		}
	}

	// Update ticket status
	ticket.Status = commonStructs.Ongoing

	result = db.Save(&ticket)

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Error While Updating Data From Database",
			Data:    nil,
		}
	}

	// TODO: Create invoice by calling payment service

	return utils.ResponseBody{
		Code:    fiber.StatusOK,
		Message: "Ticket Booked Successfully",
		Data:    fiber.Map{"status": "ongoing"},
	}
}
