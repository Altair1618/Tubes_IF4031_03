package ticketService

import (
	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateStatusService(payload commonStructs.TicketUpdateStatusServicePayload) utils.ResponseBody {

	db, _ := configs.GetGormClient()

	var ticketInvoiceBooking models.TicketInvoiceBooking

	// Change all ticket status to booked
	db.Where("invoice_id = ?", payload.InvoiceId).First(&ticketInvoiceBooking)

	var ticket models.Ticket
	db.Where("id = ?", ticketInvoiceBooking.TicketId).First(&ticket)

	if payload.Status == "FAILED" {
		utils.GeneratePDF(false, ticketInvoiceBooking.BookingId.String(), commonStructs.FailedPDFPayload{
			ErrorMessage: "Payment process failed",
		})

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to update ticket status",
		}
	}

	db.Transaction(func(tx *gorm.DB) error {
		ticket.Status = "BOOKED"
		tx.Save(&ticket)

		// Generate PDF
		utils.GeneratePDF(true, ticketInvoiceBooking.BookingId.String(), commonStructs.SuccessPDFPayload{
			Price: ticket.Price,
			Seat:  ticket.SeatId,
		})

		// TODO: Sent pdf to client service
		return nil
	})

	return utils.ResponseBody{
		Code:    200,
		Message: "Ticket status successfully updated",
	}
}
