package main

import (
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/routes"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.Bootstrap()
	routes.Routes(app)
	utils.Serve(app)
}
