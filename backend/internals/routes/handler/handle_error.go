package handler

import (
	"backend/internals/entities/response"
	"encoding/json"
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

	// Case of *response.GenericError
	if e, ok := err.(*response.GenericError); ok {
		// Check specific error types within the GenericError
		if e.Err != nil {
			if syntaxErr, ok := e.Err.(*json.SyntaxError); ok {
				return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
					Success: false,
					Code:    "JSON_SYNTAX_ERROR",
					Message: "Invalid JSON format in the request body",
					Error:   syntaxErr.Error(),
				})
			}
			if unmarshalErr, ok := e.Err.(*json.UnmarshalTypeError); ok {
				return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
					Success: false,
					Code:    "JSON_UNMARSHAL_TYPE_ERROR",
					Message: "JSON type mismatch",
					Error:   unmarshalErr.Error(),
				})
			}
			// Handle other errors, like validation errors
			if validationErrs, ok := e.Err.(validator.ValidationErrors); ok {
				return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
					Success: false,
					Code:    "VALIDATION_FAILED",
					Message: "Request body validation failed",
					Error:   validationErrs.Error(),
				})
			}

			// Default error handling if no specific error type matches
			return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				Error:   e.Err.Error(),
			})
		}
	}

	// Default generic error response for unknown errors
	return c.Status(fiber.StatusInternalServerError).JSON(
		response.ErrorResponse{
			Success: false,
			Code:    "UNKNOWN_SERVER_SIDE_ERROR",
			Message: "Unknown server side error",
			Error:   err.Error(),
		},
	)
}
