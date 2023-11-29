package routes

import (
	testController "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/controllers/test"
	"github.com/gofiber/fiber/v2"
)

func testRouteV1(v1 fiber.Router) {
	test := v1.Group("/test")

	test.Get("/probability", testController.TestProbabilityController)
}