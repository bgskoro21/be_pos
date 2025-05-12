package routes

import (
	controller "bgskoro21/be-pos/controller/user"

	"github.com/gofiber/fiber/v2"
)

type RoutConfig struct{
	UserController controller.UserController
}

func SetupRoutes(app *fiber.App, cfg RoutConfig){
	api := app.Group("/api/v1")

	UserRoutes(api, cfg.UserController)
}