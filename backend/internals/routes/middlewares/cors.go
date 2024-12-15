package middlewares

import (
	"backend/internals/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors = func() fiber.Handler {
	origins := ""
	for i, s := range config.Env.ServerOrigins {
		origins += *s
		if i < len(config.Env.ServerOrigins)-1 {
			origins += ", "
		}
	}
	config := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(config)
}()
