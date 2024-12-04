package handler

import (
	"backend/internals/entities/response"
	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: false,
		Message: "Not found",
		Error:   "404_NOT_FOUND",
	})
}
