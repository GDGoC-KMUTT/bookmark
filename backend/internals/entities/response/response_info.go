package response

import (
	"github.com/gofiber/fiber/v2"
)

type InfoResponse[T any] struct {
	Success bool   `json:"success"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func Ok[T any](ctx *fiber.Ctx, data T, meta ...any) error {
	if len(meta) > 0 {
		return ctx.Status(200).JSON(InfoResponse[T]{
			Code: fiber.StatusOK,
			Data: data,
		})
	}

	return ctx.Status(200).JSON(InfoResponse[T]{
		Code: fiber.StatusOK,
		Data: data,
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
