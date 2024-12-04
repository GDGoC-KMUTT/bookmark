package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

var Recover = func() fiber.Handler {
	if viper.GetInt("environment") == 1 {
		return func(c *fiber.Ctx) error {
			fmt.Println("[CALL] " + c.Method() + " " + c.Path() + " " + string(c.Request().Body()))
			return c.Next()
		}
	}
	return recover2.New()
}
