package routes

import (
	controller "bgskoro21/be-pos/controller/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userController controller.UserController){
	userRoutes := router.Group("/users")

	userRoutes.Post("/register", userController.Create)

}