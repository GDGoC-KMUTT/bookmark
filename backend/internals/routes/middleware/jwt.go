package middleware

import (
	"backend/internals/config"
	"backend/internals/entities/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var Jwt = func() fiber.Handler {
	config := jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(*config.Env.SecretKey)},
		TokenLookup: "cookie:login",
		ContextKey:  "userId",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(response.ErrorResponse{
				Code:    strconv.Itoa(fiber.StatusUnauthorized),
				Message: "Unauthorized access",
			})
		},
	}
	return jwtware.New(config)
}
