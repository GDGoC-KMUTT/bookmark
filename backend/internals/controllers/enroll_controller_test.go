package controllers_test

import (
	"backend/internals/controllers"
	mockServices "backend/mocks/services"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestEnrollController(mockEnrollService *mockServices.EnrollServices) *fiber.App {
	app := fiber.New()

	// Initialize the controller
	controller := controllers.NewEnrollController(mockEnrollService)

	// Middleware to simulate JWT Locals
	app.Use(func(c *fiber.Ctx) error {
		// Simulate a missing JWT for unauthorized requests
		return c.Next()
	})

	// Register the route
	app.Post("/enroll/:courseId", controller.EnrollInCourse)
	return app
}

func TestEnrollInCourseUnauthorized(t *testing.T) {
	is := assert.New(t)

	// Arrange
	mockEnrollService := new(mockServices.EnrollServices)
	app := setupTestEnrollController(mockEnrollService)

	mockCourseId := "1"

	// Create a request without a valid token (simulate unauthorized access)
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/enroll/%s", mockCourseId), nil)

	// Test
	res, err := app.Test(req, -1) // No timeout

	// Assert
	is.Nil(err)
	is.Equal(http.StatusUnauthorized, res.StatusCode)

	// Ensure the EnrollUser method is not called
	mockEnrollService.AssertNotCalled(t, "EnrollUser", mock.Anything, mock.Anything)
}
