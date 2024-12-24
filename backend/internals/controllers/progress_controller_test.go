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

	app.Get("/progress/:courseID/percentage", progressController.GetCompletionPercentage)
	return app
}

func (suite *ProgressControllerTestSuite) TestGetCompletionPercentageWhenSuccess() {
	is := assert.New(suite.T())

	mockProgressService := new(mockServices.ProgressService)

	app := setupTestProgressController(mockProgressService)

	mockCourseID := uint(10)
	expectedPercentage := 75.0

	mockProgressService.EXPECT().GetCompletionPercentage(mock.Anything, mockCourseID).Return(expectedPercentage, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/progress/%d/percentage", mockCourseID), nil)
	res, err := app.Test(req)

	var responsePayload map[string]float64
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(expectedPercentage, responsePayload["completion_percentage"])
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
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Invalid courseID", errResponse["error"])
}

func TestProgressController(t *testing.T) {
	suite.Run(t, new(ProgressControllerTestSuite))
}
