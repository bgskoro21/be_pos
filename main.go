package main

import (
	"bgskoro21/be-pos/app"
	controller "bgskoro21/be-pos/controller/user"
	"bgskoro21/be-pos/exception"
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/pkg/logger"
	repository "bgskoro21/be-pos/repository/user"
	"bgskoro21/be-pos/routes"
	service "bgskoro21/be-pos/service/user"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main(){
	logger.SetupLogger()

	db := app.InitDB()

	appFiber := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		Prefork: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error{
			return exception.ErrorHandler(ctx, err)
		},
	})

	appFiber.Use(recover.New())

	db.AutoMigrate(
		&domain.User{},
	)

	userRepo := repository.NewUserRepository(db);
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	
	routes.SetupRoutes(appFiber, routes.RoutConfig{
		UserController: userController,
	})
	
	logger.Log.Info("Server Started")

	err := appFiber.Listen(os.Getenv("APP_PORT"))

	helper.PanicIfError(err)

}