package main

import (
	"github.com/Altair1618/IF4031_03_Payment/payment_service.git/app/configs"
	"github.com/Altair1618/IF4031_03_Payment/payment_service.git/app/routes"
	"github.com/Altair1618/IF4031_03_Payment/payment_service.git/app/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.Bootstrap()
	routes.Routes(app)
	utils.Serve(app)
}
