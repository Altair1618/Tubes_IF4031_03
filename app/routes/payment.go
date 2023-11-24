package routes

import (
	paymentController "github.com/Altair1618/IF4031_03_Payment/payment_service.git/app/controllers/payment"
	"github.com/Altair1618/IF4031_03_Payment/payment_service.git/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func paymentRouteV1(v1 fiber.Router) {
	payment := v1.Group("/payment")

	payment.Get("/", middlewares.AuthMiddleware, paymentController.GetPaymentController)
}
