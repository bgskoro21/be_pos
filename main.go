package main

import (
	"bgskoro21/be-pos/app"
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/pkg/logger"
	"bgskoro21/be-pos/routes"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main(){
	logger.SetupLogger()

	db := app.InitDB()

	appFiber := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		Prefork: true,
	})

	db.AutoMigrate()
	
	routes.SetupRoutes(appFiber)
	
	logger.Log.Info("Server Started")

	err := appFiber.Listen(os.Getenv("APP_PORT"))

	helper.PanicIfError(err)

}