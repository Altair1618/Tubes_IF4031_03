package testController

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func TestProbabilityController(c *fiber.Ctx) error {
	probability := utils.SimulateProbability(20)

	s := "Enc*de Me Plea$e"
	fmt.Println(url.PathEscape(s))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(s)))

	return utils.CreateResponseBody(c, utils.ResponseBody{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    fiber.Map{"probability": fmt.Sprintf("%t", probability)},
	})
}