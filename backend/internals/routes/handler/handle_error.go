package handler

import (
	"backend/internals/entities/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Case of *fiber.Error.
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(response.ErrorResponse{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(e.Error()), " ", "_"),
			Message: e.Error(),
			Error:   e.Error(),
		})
	}

	if e, ok := err.(*response.GenericError); ok {
		if len(e.Code) == 0 {
			e.Code = "GENERIC_ERROR"
		}

		if e.Err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				Error:   e.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Success: false,
			Code:    e.Code,
			Message: e.Message,
			Error:   e.Error(),
		})
	}

	// Case of validator.ValidationErrors
	if e, ok := err.(validator.ValidationErrors); ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Code:    "VALIDATION_FAILED",
			Message: "Information validation failed",
			Error:   e.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(
		response.ErrorResponse{
			Success: false,
			Code:    "UNKNOWN_SERVER_SIDE_ERROR",
			Message: "Unknown server side error",
			Error:   err.Error(),
		},
	)
}
