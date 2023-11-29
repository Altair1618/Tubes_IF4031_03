package ticketService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func UpdateStatusService(payload commonStructs.UpdateTicketStatusServicePayload) utils.ResponseBody {

	db, _ := configs.GetGormClient()

	var ticketInvoiceBooking models.TicketInvoiceBooking

	// Change all ticket status to booked
	db.Where("invoice_id = ?", payload.InvoiceId).First(&ticketInvoiceBooking)

	var ticket models.Ticket
	db.Where("id = ?", ticketInvoiceBooking.TicketId).First(&ticket)

	if payload.Status == "FAILED" {
		url, err := utils.GeneratePDF(false, payload.UserId, ticketInvoiceBooking.BookingId.String(), commonStructs.FailedPDFPayload{
			ErrorMessage: payload.Message,
		})

		if err != nil {
			return utils.ResponseBody{
				Code:    fiber.StatusInternalServerError,
				Message: "something went wrong while generating pdf report",
			}
		}

		// TODO: Call webhook on client service containing the url and status
		fmt.Println(url)

		return utils.ResponseBody{
			Code:    fiber.StatusOK,
			Message: "ticket status sucessfully updated",
		}
	}

	// begin transaction
	tx := db.Begin()

	ticket.Status = "BOOKED"
	tx.Save(&ticket)

	// Generate PDF
	url, err := utils.GeneratePDF(true, payload.UserId, ticketInvoiceBooking.BookingId.String(), commonStructs.SuccessPDFPayload{
		Price: ticket.Price,
		Seat:  ticket.SeatId,
	})

	// TODO: Call webhook on client service containing the url and status
	fmt.Println(url)

	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return utils.ResponseBody{
			Code:    500,
			Message: fmt.Sprintf("failed to generate pdf: %s", err.Error()),
		}
	}

	// TODO: Sent pdf to client service

	tx.Commit()

	// end transaction

	return utils.ResponseBody{
		Code:    200,
		Message: "ticket status successfully updated",
	}
}
