package routes

import (
	controller "bgskoro21/be-pos/controller/user"
	"bgskoro21/be-pos/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userController controller.UserController){
	userRoutes := router.Group("/users")

	userRoutes.Post("/register", userController.Create)
	userRoutes.Post("/login", userController.Login)
	userRoutes.Post("/refresh", userController.Refresh)

	private := userRoutes.Group("/", middleware.JWTMiddleware())
	private.Get("/profile", userController.FindById)
}