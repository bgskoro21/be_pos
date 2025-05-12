package exception

import (
	"bgskoro21/be-pos/helper"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Validator error
	if ve, ok := err.(helper.ValidationError); ok {
		
		return helper.SendResponse(ctx, fiber.StatusBadRequest, nil, ve)
	}

	// Not found error (custom)
	if nf, ok := err.(NotFoundError); ok {
		return helper.SendResponse(ctx, fiber.StatusNotFound, nil, nf.Error)
	}

	// Conflict Error
	if cf, ok := err.(ConflictError); ok{
		return helper.SendResponse(ctx, fiber.StatusConflict, nil, cf.Error())
	}

	// Default: Internal server error
	return helper.SendResponse(ctx, fiber.StatusInternalServerError, nil, err)
}