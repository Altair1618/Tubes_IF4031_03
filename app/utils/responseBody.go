package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseBody struct {
	Code    int ``
	Message string
	Data    fiber.Map
}

func CreateResponseBody(c *fiber.Ctx, responseBody ResponseBody) error {
	return c.Status(responseBody.Code).JSON(fiber.Map{
		"code":    responseBody.Code,
		"message": responseBody.Message,
		"data":    responseBody.Data,
	})
}
