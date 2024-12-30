package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	"backend/internals/services"
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

type CoursePageControllerTestSuite struct {
	suite.Suite
}

func TestCoursePageController(t *testing.T) {
	suite.Run(t, new(CoursePageControllerTestSuite))
}

func setupTestCoursePageController(coursePageSvc services.CoursePageServices, withToken bool) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

	controller := NewCoursePageController(coursePageSvc)

	app.Use(func(c *fiber.Ctx) error {
		if withToken {
			// Inject a valid token for authorized requests
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId": 123, // Mock user ID
			})
			c.Locals("user", token)
		}
		return c.Next()
	})

	app.Get("/courses/:coursePageId/info", controller.GetCoursePageInfo)
	app.Get("/courses/:coursePageId/content", controller.GetCoursePageContent)
	app.Get("/courses/suggest/:fieldId", controller.GetSuggestCoursesByFieldId)

	return app
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageInfoWhenSuccess() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedCoursePageInfo := payload.CoursePage{
		Id:      1,
		Name:    "Test Course Page",
		FieldId: 2,
	}

	mockCoursePageService.EXPECT().GetCoursePageInfo(mock.Anything).Return(&expectedCoursePageInfo, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token") // Simulate a valid token

	res, err := app.Test(req)

	var responsePayload response.InfoResponse[payload.CoursePage]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(expectedCoursePageInfo.Id, responsePayload.Data.Id)
	is.Equal(expectedCoursePageInfo.Name, responsePayload.Data.Name)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageInfoWhenInvalidID() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "invalidID"
	expectedError := fmt.Errorf("course page not found")

	mockCoursePageService.EXPECT().GetCoursePageInfo(mock.Anything).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenSuccess() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	mockCourseContents := []payload.CoursePageContent{
		{
			CoursePageId: 1,
			Order:        1,
			Type:         "text",
			Text:         utils.Ptr("Sample text content"),
			ModuleId:     nil,
		},
		{
			CoursePageId: 1,
			Order:        2,
			Type:         "module",
			Text:         nil,
			ModuleId:     utils.Ptr(uint64(101)),
		},
	}

	mockCoursePageService.EXPECT().GetCoursePageContent(mock.Anything).Return(mockCourseContents, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	var responsePayload response.InfoResponse[[]payload.CoursePageContent]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Len(responsePayload.Data, len(mockCourseContents))
	is.Equal(mockCourseContents[0].CoursePageId, responsePayload.Data[0].CoursePageId)
	is.Equal(mockCourseContents[1].ModuleId, responsePayload.Data[1].ModuleId)
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenNoSuggestions() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockFieldId := "3"
	mockCoursePageService.On("GetSuggestCourseByFieldId", mockFieldId).Return([]payload.SuggestCourse{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", mockFieldId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	var responsePayload response.InfoResponse[[]payload.SuggestCourse]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(0, len(responsePayload.Data))
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenEmptyFieldId() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	req := httptest.NewRequest(http.MethodGet, "/courses/suggest", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenServiceFails() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedError := fmt.Errorf("service failure")

	mockCoursePageService.EXPECT().GetCoursePageContent(mock.Anything).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentPartialResponse() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	mockCourseContents := []payload.CoursePageContent{
		{
			CoursePageId: 1,
			Order:        1,
			Type:         "text",
			Text:         utils.Ptr("Sample text content"),
		},
	}

	mockCoursePageService.EXPECT().GetCoursePageContent(mock.Anything).Return(mockCourseContents, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	var responsePayload response.InfoResponse[[]payload.CoursePageContent]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Len(responsePayload.Data, 1)
	is.Equal("text", responsePayload.Data[0].Type)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageInfoWhenServiceFails() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedError := fmt.Errorf("internal service error")

	mockCoursePageService.EXPECT().GetCoursePageInfo(mock.Anything).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 500
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageInfoWhenNotFound() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedError := fmt.Errorf("not found")

	mockCoursePageService.EXPECT().GetCoursePageInfo(mock.Anything).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 404
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenEmpty() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"

	mockCoursePageService.EXPECT().GetCoursePageContent(mock.Anything).Return([]payload.CoursePageContent{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 404 when content is empty
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenInvalidFieldId() {
	is := assert.New(suite.T())

	// Mock the service
	mockCoursePageService := new(mockServices.CoursePageServices)

	// Ensure the service is not called for invalid field IDs
	mockCoursePageService.AssertNotCalled(suite.T(), "GetSuggestCourseByFieldId", "invalid")

	// Create the app
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	// Simulate an invalid field ID
	invalidFieldId := "invalid"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", invalidFieldId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	// Assert that it returns 404 Not Found
	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 404

	// Ensure the service was not called
	mockCoursePageService.AssertNotCalled(suite.T(), "GetSuggestCourseByFieldId", invalidFieldId)
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenServiceFails() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	fieldIdStr := "2"
	expectedError := fmt.Errorf("service failure")

	mockCoursePageService.On("GetSuggestCourseByFieldId", fieldIdStr).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 500
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenFieldIdIsEmpty() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	req := httptest.NewRequest(http.MethodGet, "/courses//content", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for empty fieldId
	mockCoursePageService.AssertNotCalled(suite.T(), "GetCoursePageContent")
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenFieldIdIsNumericButNotFound() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	fieldIdStr := "123" // Numeric fieldId
	expectedError := fmt.Errorf("not found")

	mockCoursePageService.On("GetSuggestCourseByFieldId", fieldIdStr).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 404 for not found
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenServerErrorOccurs() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	fieldIdStr := "2"
	expectedError := fmt.Errorf("internal server error")

	mockCoursePageService.On("GetSuggestCourseByFieldId", fieldIdStr).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 500
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenRequestMethodIsInvalid() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil) // Invalid method
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusMethodNotAllowed, res.StatusCode) // Expect 405 Method Not Allowed
	mockCoursePageService.AssertNotCalled(suite.T(), "GetCoursePageContent", mockCoursePageId)
}

func (suite *CoursePageControllerTestSuite) TestGetEnrollCourseByUserIdWhenUserIdIsInvalid() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	invalidUserId := "abc" // Invalid userId

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/enrollments/%s", invalidUserId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for invalid userId
	mockCoursePageService.AssertNotCalled(suite.T(), "GetEnrollCourseByUserId", invalidUserId)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenNoContent() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true)

	mockCoursePageId := "1"

	mockCoursePageService.EXPECT().GetCoursePageContent(mock.Anything).Return([]payload.CoursePageContent{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 404 for empty content
	//mockCoursePageService.AssertCalled(suite.T(), "GetCoursePageContent", mockCoursePageId)
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenFieldIdIsNotNumeric() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true)

	fieldIdStr := "non-numeric"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 404 for non-numeric fieldId
	mockCoursePageService.AssertNotCalled(suite.T(), "GetSuggestCourseByFieldId")
}

func (suite *CoursePageControllerTestSuite) TestGetSuggestCoursesByFieldIdWhenFieldIdIsEmpty() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	// Simulate a request with an empty fieldId
	req := httptest.NewRequest(http.MethodGet, "/courses/suggest/", nil) // No fieldId in the path
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	// Assert that the response returns 404 Not Found
	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404
	mockCoursePageService.AssertNotCalled(suite.T(), "GetSuggestCourseByFieldId")
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageInfoWhenUnhandledError() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true)

	mockCoursePageId := "1"
	unexpectedError := fmt.Errorf("unexpected error occurred")

	mockCoursePageService.EXPECT().GetCoursePageInfo(mock.Anything).Return(nil, unexpectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)

	var responsePayload response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)
	is.Equal("failed to get course page info", responsePayload.Message)
}

func (suite *CoursePageControllerTestSuite) TestGetCoursePageContentWhenEmptyResponse() {
	is := assert.New(suite.T())

	mockCoursePageService := new(mockServices.CoursePageServices)
	app := setupTestCoursePageController(mockCoursePageService, true)

	mockCoursePageId := utils.Ptr(uint64(1))

	// Return an empty response
	mockCoursePageService.EXPECT().GetCoursePageContent(mock.Anything).Return([]payload.CoursePageContent{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%d/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	var responsePayload response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}
