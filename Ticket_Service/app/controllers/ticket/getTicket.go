package ticketController

import (
	"time"

	"github.com/Altair1618/IF4031_03_Ticket/app/configs"
	"github.com/Altair1618/IF4031_03_Ticket/app/models"
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
