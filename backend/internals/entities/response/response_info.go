package response

import (
	"github.com/gofiber/fiber/v2"
)

type InfoResponse[T any] struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data"`
}

func Ok[T any](ctx *fiber.Ctx, data T, meta ...any) error {
	if len(meta) > 0 {
		return ctx.Status(200).JSON(InfoResponse[T]{
			Success: true,
			Code:    fiber.StatusOK,
			Data:    data,
		})
	}

	return ctx.Status(200).JSON(InfoResponse[T]{
		Success: true,
		Code:    fiber.StatusOK,
		Data:    data,
	})
}

func Created[T any](ctx *fiber.Ctx, data T, meta ...any) error {
	if len(meta) > 0 {
		return ctx.Status(200).JSON(InfoResponse[T]{
			Code: fiber.StatusCreated,
			Data: data,
		})
	}

	return ctx.Status(200).JSON(InfoResponse[T]{
		Code: fiber.StatusCreated,
		Data: data,
	})
}
