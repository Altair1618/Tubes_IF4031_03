package routes

import (
	paymentController "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/controllers/payment"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func paymentRouteV1(v1 fiber.Router) {
	payment := v1.Group("/payment")

	payment.Patch("/:paymentToken", middlewares.AuthMiddleware, paymentController.ProcesPaymentController)
}
