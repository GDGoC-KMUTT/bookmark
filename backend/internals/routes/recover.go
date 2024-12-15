package routes

import (
	"backend/internals/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

var Recover = func() fiber.Handler {
	if *config.Env.Environment == 0 {
		return func(c *fiber.Ctx) error {
			fmt.Println("[CALL] " + c.Method() + " " + c.Path() + " " + string(c.Request().Body()))
			return c.Next()
		}
	}
	return recover2.New()
}
