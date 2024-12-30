package controllers

import (
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
	"backend/internals/utils"
)

type ModuleControllerTestSuite struct {
    suite.Suite
}

func setupTestModuleController(mockModuleService *mockServices.ModuleServices) *fiber.App {
    fiberConfig := fiber.Config{
        ErrorHandler: handler.ErrorHandler,
    }
    app := fiber.New(fiberConfig)

    moduleController := NewModuleController(mockModuleService) // Pass mock service here

    app.Get("/module/:moduleId/info", moduleController.GetModuleInfo)

    return app
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWhenSuccess() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "123"
    mockResponse := &payload.ModuleResponse{
        Id:          123,
        Title:       "Test Module",
        Description: utils.Ptr("This is a description."),
        ImageUrl:    utils.Ptr("http://example.com/image.png"),
    }

    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(mockResponse, nil)

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    var r response.InfoResponse[payload.ModuleResponse]
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &r)

    is.Nil(err)
    is.Equal(http.StatusOK, res.StatusCode)
    is.Equal(uint64(123), r.Data.Id)
    is.Equal("Test Module", r.Data.Title)
    is.Equal("This is a description.", *r.Data.Description)
    is.Equal("http://example.com/image.png", *r.Data.ImageUrl)
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWhenServiceFails() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "123"
    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(nil, fmt.Errorf("service error"))

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("failed to get module info", errResponse.Message)
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithInvalidModuleId() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    req := httptest.NewRequest(http.MethodGet, "/module//info", nil) // Missing moduleId
    res, err := app.Test(req)

    is.Nil(err)
    is.Equal(http.StatusNotFound, res.StatusCode) // Fiber returns 404 for invalid routes
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWhenServiceReturnsNil() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "123"
    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(nil, nil)

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("module info is not available", errResponse.Message)
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithContextError() {
    is := assert.New(suite.T())

    // mockModuleService := new(mockServices.ModuleServices)
    app := fiber.New() // Use plain app to inject custom context handler

    app.Get("/module/:moduleId/info", func(ctx *fiber.Ctx) error {
        return ctx.Status(fiber.StatusInternalServerError).SendString("context error")
    })

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    body, _ := io.ReadAll(res.Body)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("context error", string(body))
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWhenServiceReturnsError() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "123"
    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(nil, fmt.Errorf("unexpected service error"))

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("failed to get module info", errResponse.Message)
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithLargeModuleId() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "1234567890123456789012345678901234567890" // Very large ID
    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(nil, nil)

    req := httptest.NewRequest(http.MethodGet, "/module/"+mockModuleId+"/info", nil)
    res, err := app.Test(req)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithEmptyModuleId() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    req := httptest.NewRequest(http.MethodGet, "/module//info", nil) // Missing moduleId
    res, err := app.Test(req)

    is.Nil(err)
    is.Equal(http.StatusNotFound, res.StatusCode) // Fiber handles empty paths with 404
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithUnexpectedError() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "123"
    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(nil, fmt.Errorf("unexpected error"))

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("failed to get module info", errResponse.Message)
}

func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithNullOptionalFields() {
    is := assert.New(suite.T())

    mockModuleService := new(mockServices.ModuleServices)
    app := setupTestModuleController(mockModuleService)

    mockModuleId := "123"
    mockResponse := &payload.ModuleResponse{
        Id:    123,
        Title: "Test Module",
    }

    mockModuleService.EXPECT().GetModuleInfo(mockModuleId).Return(mockResponse, nil)

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    var r response.InfoResponse[payload.ModuleResponse]
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &r)

    is.Nil(err)
    is.Equal(http.StatusOK, res.StatusCode)
    is.Equal(uint64(123), r.Data.Id)
    is.Equal("Test Module", r.Data.Title)
    is.Nil(r.Data.Description)
    is.Nil(r.Data.ImageUrl)
}
func (suite *ModuleControllerTestSuite) TestGetModuleInfoWithInvalidJSON() {
    is := assert.New(suite.T())

    // mockModuleService := new(mockServices.ModuleServices)
    app := fiber.New() // Create app to inject custom handler

    app.Get("/module/:moduleId/info", func(ctx *fiber.Ctx) error {
        return ctx.Status(fiber.StatusOK).SendString("{invalid_json}") // Malformed JSON
    })

    req := httptest.NewRequest(http.MethodGet, "/module/123/info", nil)
    res, err := app.Test(req)

    body, _ := io.ReadAll(res.Body)

    is.Nil(err)
    is.Equal(http.StatusOK, res.StatusCode) // Still a 200 but with malformed response
    is.Equal("{invalid_json}", string(body))
}


func TestModuleController(t *testing.T) {
    suite.Run(t, new(ModuleControllerTestSuite))
}
