package routes

import (
	"bgskoro21/be-pos/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/", func(c *fiber.Ctx) error {
		logger.Log.Info("Send Hello World")
		return c.SendString("Hello, World!")
	})
}