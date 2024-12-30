package controllers_test

import (
	"backend/internals/controllers"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	mockServices "backend/mocks/services"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
)

type ModuleStepControllerTestSuite struct {
	suite.Suite
}

func setupTestModuleStepController(mockModuleStepService *mockServices.ModuleStepServices) *fiber.App {
	app := fiber.New()

	// Initialize the controller
	moduleStepController := controllers.NewModuleStepController(mockModuleStepService)

	// Middleware to simulate JWT Locals
	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{"userId": float64(123)} // Simulate a valid userId claim
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	// Register the route
	app.Get("/step/:moduleId/info", moduleStepController.GetModuleSteps)
	return app
}

func (suite *ModuleStepControllerTestSuite) TestGetModuleStepsWhenSuccess() {
    is := assert.New(suite.T())

    mockModuleStepService := new(mockServices.ModuleStepServices)
    app := setupTestModuleStepController(mockModuleStepService)

    moduleID := "module123"
    mockSteps := []payload.ModuleStepResponse{
        {Id: 1, Title: "Step 1", Check: true},
    }

    mockModuleStepService.EXPECT().GetModuleSteps(uint(123), moduleID).Return(mockSteps, nil)

    req := httptest.NewRequest(http.MethodGet, "/step/module123/info", nil)
    req.Header.Set("Authorization", "Bearer mockToken")

    res, err := app.Test(req)

    var responsePayload response.InfoResponse[[]payload.ModuleStepResponse]
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &responsePayload)

    is.Nil(err)
    is.Equal(http.StatusOK, res.StatusCode)
    is.Equal(uint64(1), responsePayload.Data[0].Id) // Updated to match the actual type
    is.Equal("Step 1", responsePayload.Data[0].Title)
    is.True(responsePayload.Data[0].Check)
}

func (suite *ModuleStepControllerTestSuite) TestGetModuleStepsWhenServiceFails() {
    is := assert.New(suite.T())

    mockModuleStepService := new(mockServices.ModuleStepServices)
    app := setupTestModuleStepController(mockModuleStepService)

    moduleID := "module123"

    mockModuleStepService.EXPECT().GetModuleSteps(uint(123), moduleID).Return(nil, errors.New("service error"))

    req := httptest.NewRequest(http.MethodGet, "/step/module123/info", nil)
    req.Header.Set("Authorization", "Bearer mockToken")

    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("service error", errResponse.Message)
}

func (suite *ModuleStepControllerTestSuite) TestGetModuleStepsWhenModuleIDMissing() {
    is := assert.New(suite.T())

    mockModuleStepService := new(mockServices.ModuleStepServices)
    app := setupTestModuleStepController(mockModuleStepService)

    req := httptest.NewRequest(http.MethodGet, "/step//info", nil)
    req.Header.Set("Authorization", "Bearer mockToken")

    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusBadRequest, res.StatusCode)
    is.Equal("missing module ID", errResponse.Message)
}
