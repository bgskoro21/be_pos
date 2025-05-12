package helper

import (
	"bgskoro21/be-pos/model/dto"

	"github.com/gofiber/fiber/v2"
)

func SendResponse(ctx *fiber.Ctx, statusCode int, data interface{}, errors interface{}) error {
	response := dto.ApiResponse{
		Code:   statusCode,
		Data:   data,
		Errors: errors,
	}
	 
	return ctx.Status(statusCode).JSON(response)
}