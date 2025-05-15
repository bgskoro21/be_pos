package controller

import (
	"github.com/gofiber/fiber/v2"
)

type UserController interface{
	Create(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
}