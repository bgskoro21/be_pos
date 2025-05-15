package controller

import (
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/model/dto"
	service "bgskoro21/be-pos/service/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct{
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController{
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(ctx *fiber.Ctx) error{
	var request dto.RegisterUserRequest

	if err := ctx.BodyParser(&request); err != nil{
		helper.PanicIfError(err)
	}

	user, err := controller.UserService.Register(request);

	helper.PanicIfError(err)

	return helper.SendResponse(ctx, fiber.StatusCreated, user, nil)
}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error{
	var request dto.LoginRequest

	if err := ctx.BodyParser(&request); err != nil{
		helper.PanicIfError(err)
	}

	request.UserAgent = string(ctx.Request().Header.UserAgent())
	request.IPAddress = ctx.IP()

	data, err := controller.UserService.Login(request);

	helper.PanicIfError(err)

	return helper.SendResponse(ctx, fiber.StatusOK, data, nil)
}

func (controller *UserControllerImpl) Refresh(ctx *fiber.Ctx) error{
	var request dto.RefreshTokenRequest

	if err := ctx.BodyParser(&request); err != nil{
		helper.PanicIfError(err)
	}

	request.UserAgent = string(ctx.Request().Header.UserAgent())
	request.IPAddress = ctx.IP()

	data, err := controller.UserService.Refresh(request)

	helper.PanicIfError(err)

	return helper.SendResponse(ctx, fiber.StatusOK, data, nil)
}

func (controller *UserControllerImpl) FindById(ctx *fiber.Ctx) error{
	val := ctx.Locals("user_id")
	userId, ok := val.(float64)

	if !ok{
		panic(fmt.Sprintf("Expected user_id to be float64, but got %T", val))
	}
	
	user, _ := controller.UserService.FindById(uint(userId))

	return helper.SendResponse(ctx, fiber.StatusOK, user, nil)
}