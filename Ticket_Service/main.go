package main

import (
	"github.com/Altair1618/IF4031_03_Ticket/app/configs"
	"github.com/Altair1618/IF4031_03_Ticket/app/routes"
	"github.com/Altair1618/IF4031_03_Ticket/app/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.Bootstrap()
	routes.Routes(app)
	utils.Serve(app)
}
