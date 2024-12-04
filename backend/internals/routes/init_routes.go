package routes

import (
	"backend/internals/config"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

func SetupRoutes() {

	serverAddr := fmt.Sprintf("%s:%d", config.Env.ServerHost, config.Env.ServerPort)

	// Initialize fiber instance
	app := NewFiberApp()

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.InfoResponse{
			Success: true,
			Message: "BOOKMARK_API_ROOT",
		})
	})

	allowOrigins := strings.Join(config.Env.ServerOrigins, ",")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
	}))

	v1 := app.Group("/api/v1")
	v1.Static("/static", "./resources/static")

	// Custom handler to set Content-Type header based on file extension
	v1.Use("/static", func(c *fiber.Ctx) error {
		filePath := c.Path()
		contentType := getContentType(filePath)
		c.Set("Content-Type", contentType)
		return c.Next()
	})

	v1.Get("/swagger/*", swagger.HandlerDefault)
	v1.Use(Recover())

	ListenAndServe(app, serverAddr)
}

func getContentType(filename string) string {
	extension := strings.ToLower(filepath.Ext(filename))
	switch extension {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	// Add more cases for other file types if needed
	default:
		return "application/octet-stream"
	}
}

func NewFiberApp() *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	return app
}

func ListenAndServe(app *fiber.App, serverAddr string) {
	err := app.Listen(serverAddr)
	if err != nil {
		panic(fmt.Errorf("[Server] Unable to start server: %w", err))
	}
	logrus.Debug("[Server] Server started successfully")
}
