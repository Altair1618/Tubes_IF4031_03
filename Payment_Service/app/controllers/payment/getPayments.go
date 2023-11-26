package paymentController

import (
	"github.com/gofiber/fiber/v2"
)

func GetPaymentController(c *fiber.Ctx) error {
	userInfo := c.Locals("userInfo").(map[string]string)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   userInfo,
	})
}
