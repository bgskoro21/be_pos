package main

import (
	"bgskoro21/be-pos/app"
	"bgskoro21/be-pos/exception"
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/pkg/logger"
	"bgskoro21/be-pos/routes"
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
		Prefork: false,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error{
			return exception.ErrorHandler(ctx, err)
		},
	})

	appFiber.Use(recover.New())

	db.AutoMigrate(
		&domain.User{},
		&domain.RefreshToken{},
	)

	container := app.InitContainer(db);
	
	routes.SetupRoutes(appFiber, routes.RoutConfig{
		UserController: container.UserController,
	})
	
	logger.Log.Info("Server Started")

	err := appFiber.Listen(":" + os.Getenv("APP_PORT"))

	helper.PanicIfError(err)

}