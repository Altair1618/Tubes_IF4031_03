package ticketService

import (
	"encoding/json"
	"fmt"
	"os"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func UpdateStatusService(payload commonStructs.UpdateTicketStatusServicePayload) utils.ResponseBody {

	db, _ := configs.GetGormClient()

	var ticketInvoiceBooking models.TicketInvoiceBooking

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
		agent := fiber.Patch(fmt.Sprintf("%s/bookings/%s", viper.Get("CLIENT_SERVICE_BASE_URL"), ticketInvoiceBooking.BookingId))
		agent.Set("Authorization", fmt.Sprintf("Bearer %s", payload.JWTToken))
		agent.JSON(fiber.Map{
			"status": payload.Status,
			"pdf":    url,
		})
		statusCode, body, errs := agent.Bytes()

		if len(errs) > 0 {
			_ = os.Remove(fmt.Sprintf("./public%s", url))
			return utils.ResponseBody{
				Code:    fiber.StatusInternalServerError,
				Message: errs[0].Error(),
			}
		}

		var webhookResponse commonStructs.HttpResponse[interface{}]
		if err := json.Unmarshal(body, &webhookResponse); err != nil {
			_ = os.Remove(fmt.Sprintf("./public%s", url))
			return utils.ResponseBody{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			}
		}

		if statusCode != fiber.StatusOK {
			_ = os.Remove(fmt.Sprintf("./public%s", url))
			return utils.ResponseBody{
				Code:    fiber.StatusInternalServerError,
				Message: webhookResponse.Message,
			}
		}

		return utils.ResponseBody{
			Code:    fiber.StatusOK,
			Message: "ticket status sucessfully updated",
		}
	}

	// begin transaction
	tx := db.Begin()

	// Change ticket status to booked
	ticket.Status = "BOOKED"
	tx.Save(&ticket)

	// Generate PDF
	url, err := utils.GeneratePDF(true, payload.UserId, ticketInvoiceBooking.BookingId.String(), commonStructs.SuccessPDFPayload{
		Price: ticket.Price,
		Seat:  ticket.SeatId,
	})

	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return utils.ResponseBody{
			Code:    500,
			Message: fmt.Sprintf("failed to generate pdf: %s", err.Error()),
		}
	}

	// TODO: Sent pdf to client service
	agent := fiber.Patch(fmt.Sprintf("%s/bookings/%s", viper.Get("CLIENT_SERVICE_BASE_URL"), ticketInvoiceBooking.BookingId))
	agent.Set("Authorization", fmt.Sprintf("Bearer %s", payload.JWTToken))
	agent.JSON(fiber.Map{
		"status": payload.Status,
		"pdf":    url,
	})
	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		_ = os.Remove(fmt.Sprintf("./public%s", url))
		tx.Rollback()
		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: errs[0].Error(),
		}
	}

	var webhookResponse commonStructs.HttpResponse[interface{}]
	if err := json.Unmarshal(body, &webhookResponse); err != nil {
		_ = os.Remove(fmt.Sprintf("./public%s", url))
		tx.Rollback()
		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if statusCode != fiber.StatusOK {
		_ = os.Remove(fmt.Sprintf("./public%s", url))
		tx.Rollback()
		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: webhookResponse.Message,
		}
	}

	tx.Commit()

	// end transaction

	return utils.ResponseBody{
		Code:    200,
		Message: "ticket status successfully updated",
	}
}
