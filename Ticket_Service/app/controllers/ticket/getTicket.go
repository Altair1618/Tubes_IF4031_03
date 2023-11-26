package ticketController

import "github.com/gofiber/fiber/v2"

func GetTicketController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   nil,
	})
}
