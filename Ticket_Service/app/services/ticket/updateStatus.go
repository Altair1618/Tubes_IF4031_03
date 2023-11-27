package ticketService

import (
	"fmt"

	commonStructs "github.com/Altair1618/IF4031_03_Ticket/app/common/structs"
	"github.com/Altair1618/IF4031_03_Ticket/app/configs"
	"github.com/Altair1618/IF4031_03_Ticket/app/models"
	"github.com/Altair1618/IF4031_03_Ticket/app/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateStatusService(payload commonStructs.UpdateStatusServicePayload) utils.ResponseBody {
	db, _ := configs.GetGormClient()

	var ticketInvoiceBookings []models.TicketInvoiceBooking
	var tickets []models.Ticket

	db.Transaction(func(tx *gorm.DB) error {
		// Change all ticket status to booked
		db.Where("invoice_id = ?", payload.InvoiceId).Find(&ticketInvoiceBookings)
		for _, ticketInvoiceBooking := range ticketInvoiceBookings {
			fmt.Println(ticketInvoiceBooking.InvoiceId)
		}
		var ticket models.Ticket
		db.Where("")
		tickets = append(tickets)
		return nil
	})

	return utils.ResponseBody{
		Code:    200,
		Message: "Ticket status successfully updated",
		Data: fiber.Map{
			"ticketInvoiceBookings": ticketInvoiceBookings,
		},
	}
}
