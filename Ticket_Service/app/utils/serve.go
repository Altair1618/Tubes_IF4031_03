package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Serve(app *fiber.App) {
	port := ":" + viper.Get("PORT").(string)
	if port == ":" {
		port = ":3069"
	}

	app.Listen(port)
}
