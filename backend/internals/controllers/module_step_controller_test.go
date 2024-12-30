package controllers_test

import (
	"backend/internals/controllers"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	mockServices "backend/mocks/services"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ModuleStepControllerTestSuite struct {
	suite.Suite
}

func setupTestModuleStepController(mockService *mockServices.ModuleStepServices) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

	controller := controllers.NewModuleStepController(mockService)

	app.Get("/step/:moduleId/info", controller.GetModuleSteps)

	return app
}

func (suite *ModuleStepControllerTestSuite) TestGetModuleStepsWhenSuccess() {
	is := assert.New(suite.T())

	mockService := new(mockServices.ModuleStepServices)
	app := setupTestModuleStepController(mockService)

	mockModuleID := "123"
	mockSteps := []payload.ModuleStepResponse{
		{
			Id:    1,
			Title: "Step 1",
			Check: "true",
		},
		{
			Id:    2,
			Title: "Step 2",
			Check: "false",
		},
	}

	mockService.EXPECT().GetModuleSteps(mockModuleID).Return(mockSteps, nil)

	req := httptest.NewRequest(http.MethodGet, "/step/123/info", nil)
	res, err := app.Test(req)

	var r response.InfoResponse[[]payload.ModuleStepResponse]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Len(r.Data, 2)
	is.Equal(mockSteps[0].Id, r.Data[0].Id)
	is.Equal(mockSteps[1].Id, r.Data[1].Id)
}

func (suite *ModuleStepControllerTestSuite) TestGetModuleStepsWhenServiceFails() {
	is := assert.New(suite.T())

	mockService := new(mockServices.ModuleStepServices)
	app := setupTestModuleStepController(mockService)

	mockModuleID := "123"

	mockService.EXPECT().GetModuleSteps(mockModuleID).Return(nil, fmt.Errorf("service error"))

	req := httptest.NewRequest(http.MethodGet, "/step/123/info", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get module steps", errResponse.Message)
}

func (suite *ModuleStepControllerTestSuite) TestGetModuleStepsWhenParseParamFailed() {
	is := assert.New(suite.T())

	mockService := new(mockServices.ModuleStepServices)
	app := setupTestModuleStepController(mockService)

	req := httptest.NewRequest(http.MethodGet, "/step/invalid/info", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("invalid moduleId parameter", errResponse.Message)
}

func TestModuleStepController(t *testing.T) {
	suite.Run(t, new(ModuleStepControllerTestSuite))
}
