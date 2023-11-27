package ticketController

import (
	"time"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/gofiber/fiber/v2"
)

func GetTicketController(c *fiber.Ctx) error {
	db, _ := configs.GetGormClient()
	event := models.Event{
		EventName: "Konser BP",
		EventTime: time.Now().UTC(),
		Location:  "GBK",
	}

	_ = db.Create(&event)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   nil,
	})
}
