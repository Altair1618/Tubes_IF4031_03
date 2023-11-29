package routes

import (
	invoiceController "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/controllers/invoice"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func invoiceRouteV1(v1 fiber.Router) {
	invoice := v1.Group("/invoice")

	invoice.Post("/", middlewares.AuthMiddleware, invoiceController.CreateInvoiceController)
}