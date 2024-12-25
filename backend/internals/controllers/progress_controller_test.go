package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
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

type ProgressControllerTestSuite struct {
	suite.Suite
}

func setupTestProgressController(mockProgressService *mockServices.ProgressService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	progressController := NewProgressController(mockProgressService)

	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{"userId": float64(123)}
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/progress/:courseId/percentage", progressController.GetCompletionPercentage)
	return app
}

func (suite *ProgressControllerTestSuite) TestGetCompletionPercentageWhenSuccess() {
	is := assert.New(suite.T())

	mockProgressService := new(mockServices.ProgressService)

	app := setupTestProgressController(mockProgressService)

	mockCourseID := uint(10)
	expectedPercentage := 100.0

	mockProgressService.EXPECT().GetCompletionPercentage(mock.Anything, mockCourseID).Return(expectedPercentage, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/progress/%d/percentage", mockCourseID), nil)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uint(1), // Mocked user ID
	})
	req.Header.Set("Authorization", "Bearer "+token.Raw)

	res, err := app.Test(req)

	var responsePayload struct {
		Success bool    `json:"success"`
		Code    int     `json:"code"`
		Data    float64 `json:"data"`
	}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.True(responsePayload.Success)
	is.Equal(200, responsePayload.Code)
	is.Equal(expectedPercentage, responsePayload.Data) // The completion percentage is in the `data` field
}

func (suite *ProgressControllerTestSuite) TestGetCompletionPercentageWhenFailedToGetCompletionPercentage() {
	is := assert.New(suite.T())

	mockProgressService := new(mockServices.ProgressService)

	app := setupTestProgressController(mockProgressService)

	mockCourseID := uint(10)

	mockProgressService.EXPECT().GetCompletionPercentage(mock.Anything, mockCourseID).Return(0.0, fmt.Errorf("failed to get completion percentage"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/progress/%d/percentage", mockCourseID), nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get completion percentage", errResponse.Message)
}

func (suite *ProgressControllerTestSuite) TestGetCompletionPercentageWhenInvalidCourseID() {
	is := assert.New(suite.T())

	mockProgressService := new(mockServices.ProgressService)

	app := setupTestProgressController(mockProgressService)

	req := httptest.NewRequest(http.MethodGet, "/progress/invalidCourseID/percentage", nil)

	res, err := app.Test(req)

	var errResponse map[string]string
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal("failed to decode: schema: error converting value for \"courseId\"", errResponse["error"])
}

func TestProgressController(t *testing.T) {
	suite.Run(t, new(ProgressControllerTestSuite))
}
