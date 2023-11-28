package testController

import (
	"fmt"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func TestProbabilityController(c *fiber.Ctx) error {
	probability := utils.SimulateProbability(20)
	return utils.CreateResponseBody(c, utils.ResponseBody{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    fiber.Map{"probability": fmt.Sprintf("%t", probability)},
	})
}