package controllers_test

import (
	"backend/internals/controllers"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	"backend/internals/utils"
	mockServices "backend/mocks/services"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"
)

type EnrollControllerTestSuite struct {
	suite.Suite
}

func setupTestEnrollController(mockEnrollService *mockServices.EnrollService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

	// Initialize the controller
	controller := controllers.NewEnrollController(mockEnrollService)

	// Middleware to simulate JWT Locals
	app.Use(func(c *fiber.Ctx) error {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["userId"] = float64(123)
		c.Locals("user", token)
		return c.Next()
	})

	// Register the route
	app.Post("/enroll/:courseId", controller.EnrollInCourse)
	return app
}

func TestEnrollController(t *testing.T) {
	suite.Run(t, new(EnrollControllerTestSuite))
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenSuccess() {
	is := assert.New(suite.T())

	mockEnrollService := new(mockServices.EnrollService)

	app := setupTestEnrollController(mockEnrollService)

	mockCourseId := utils.Ptr(uint64(1))

	mockEnrollService.EXPECT().EnrollUser(mock.Anything, mock.Anything).Return(nil)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/enroll/%d", mockCourseId), nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.InfoResponse[string])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenInvalidParam() {
	is := assert.New(suite.T())

	mockEnrollService := new(mockServices.EnrollService)

	app := setupTestEnrollController(mockEnrollService)

	req := httptest.NewRequest(http.MethodPost, "/enroll/invalid", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("invalid courseId parameter", r.Message)
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenInternalError() {
	is := assert.New(suite.T())

	mockEnrollService := new(mockServices.EnrollService)

	app := setupTestEnrollController(mockEnrollService)

	mockCourseId := utils.Ptr(uint64(1))

	mockEnrollService.EXPECT().EnrollUser(mock.Anything, mock.Anything).Return(fmt.Errorf("failed to enroll"))

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/enroll/%d", mockCourseId), nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to enroll course", r.Message)
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenAlreadyEnroll() {
	is := assert.New(suite.T())

	mockEnrollService := new(mockServices.EnrollService)

	app := setupTestEnrollController(mockEnrollService)

	mockCourseId := utils.Ptr(uint64(1))

	mockEnrollService.EXPECT().EnrollUser(mock.Anything, mock.Anything).Return(fmt.Errorf("user is already enrolled in this course"))

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/enroll/%d", mockCourseId), nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("user already enrolled", r.Message)
}
