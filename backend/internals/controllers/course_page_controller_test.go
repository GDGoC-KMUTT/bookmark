package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
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
	"github.com/golang-jwt/jwt/v5"
	"backend/internals/routes/handler"
)

type CoursePageControllerTestSuite struct {
	suite.Suite
}

func setupTestCoursePageController(coursePageSvc services.CoursePageServiceInterface, withToken bool) *fiber.App {
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

func TestGetCoursePageInfoWhenSuccess(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedCoursePageInfo := payload.CoursePage{
		Id:      1,
		Name:    "Test Course Page",
		FieldId: 2,
	}

	mockCoursePageService.On("GetCoursePageInfo", mockCoursePageId).Return(&expectedCoursePageInfo, nil)

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

// Helper function to return a pointer to a string
func pointerToString(s string) *string {
	return &s
}

func TestGetCoursePageInfoWhenInvalidID(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "invalidID"
	expectedError := fmt.Errorf("course page not found")

	mockCoursePageService.On("GetCoursePageInfo", mockCoursePageId).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode)
}

func TestGetCoursePageContentWhenSuccess(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	mockCourseContents := []payload.CoursePageContent{
		{
			CoursePageId: 1,
			Order:        1,
			Type:         "text",
			Text:         pointerToString("Sample text content"),
			ModuleId:     nil,
		},
		{
			CoursePageId: 1,
			Order:        2,
			Type:         "module",
			Text:         nil,
			ModuleId:     pointerToUint64(101),
		},
	}

	mockCoursePageService.On("GetCoursePageContent", mockCoursePageId).Return(mockCourseContents, nil)

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

func pointerToUint64(u uint64) *uint64 {
	return &u
}

func TestGetSuggestCoursesByFieldIdWhenNoSuggestions(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
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

func TestGetCoursePageInfoWhenNonNumericID(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "abc"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode)

	mockCoursePageService.AssertNotCalled(t, "GetCoursePageInfo", mockCoursePageId)
}

func TestGetSuggestCoursesByFieldIdWhenEmptyFieldId(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	req := httptest.NewRequest(http.MethodGet, "/courses/suggest", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode)
}

func TestGetCoursePageContentWhenServiceFails(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedError := fmt.Errorf("service failure")

	mockCoursePageService.On("GetCoursePageContent", mockCoursePageId).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func TestGetCoursePageContentPartialResponse(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	mockCourseContents := []payload.CoursePageContent{
		{
			CoursePageId: 1,
			Order:        1,
			Type:         "text",
			Text:         pointerToString("Sample text content"),
		},
	}

	mockCoursePageService.On("GetCoursePageContent", mockCoursePageId).Return(mockCourseContents, nil)

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

func TestGetCoursePageInfoWhenServiceFails(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedError := fmt.Errorf("internal service error")

	mockCoursePageService.On("GetCoursePageInfo", mockCoursePageId).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode) // Expect 500
}

func TestGetCoursePageInfoWhenNotFound(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"
	expectedError := fmt.Errorf("not found")

	mockCoursePageService.On("GetCoursePageInfo", mockCoursePageId).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404
}

func TestGetCoursePageContentWhenEmpty(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"

	mockCoursePageService.On("GetCoursePageContent", mockCoursePageId).Return([]payload.CoursePageContent{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 when content is empty
}

func TestGetSuggestCoursesByFieldIdWhenInvalidFieldId(t *testing.T) {
	is := assert.New(t)

	// Mock the service
	mockCoursePageService := new(mockServices.CoursePageServiceInterface)

	// Ensure the service is not called for invalid field IDs
	mockCoursePageService.AssertNotCalled(t, "GetSuggestCourseByFieldId", "invalid")

	// Create the app
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	// Simulate an invalid field ID
	invalidFieldId := "invalid"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", invalidFieldId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	// Assert that it returns 404 Not Found
	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404

	// Ensure the service was not called
	mockCoursePageService.AssertNotCalled(t, "GetSuggestCourseByFieldId", invalidFieldId)
}

func TestGetSuggestCoursesByFieldIdWhenServiceFails(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
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

func TestGetCoursePageContentWhenFieldIdIsEmpty(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	req := httptest.NewRequest(http.MethodGet, "/courses//content", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for empty fieldId
	mockCoursePageService.AssertNotCalled(t, "GetCoursePageContent")
}

func TestGetSuggestCoursesByFieldIdWhenNoToken(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, false) // No token

	fieldIdStr := "2"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusUnauthorized, res.StatusCode) // Expect 401
	mockCoursePageService.AssertNotCalled(t, "GetSuggestCourseByFieldId")
}

func TestGetSuggestCoursesByFieldIdWhenFieldIdIsNumericButNotFound(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	fieldIdStr := "123" // Numeric fieldId
	expectedError := fmt.Errorf("not found")

	mockCoursePageService.On("GetSuggestCourseByFieldId", fieldIdStr).Return(nil, expectedError)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for not found
}

func TestGetSuggestCoursesByFieldIdWhenServerErrorOccurs(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
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

func TestGetCoursePageInfoWhenTokenIsMalformed(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, false) // No token injection

	mockCoursePageId := "1"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer malformed-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusUnauthorized, res.StatusCode) // Expect 401 for malformed token
	mockCoursePageService.AssertNotCalled(t, "GetCoursePageInfo", mockCoursePageId)
}

func TestGetCoursePageInfoWithMissingUserIdInToken(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, false) // No valid token

	mockCoursePageId := "1"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusUnauthorized, res.StatusCode) // Expect 401 for missing userId in token
	mockCoursePageService.AssertNotCalled(t, "GetCoursePageInfo", mockCoursePageId)
}

func TestGetCoursePageInfoWhenMalformedFieldId(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "123abc" // Malformed fieldId

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for malformed fieldId
	mockCoursePageService.AssertNotCalled(t, "GetCoursePageInfo", mockCoursePageId)
}

func TestGetSuggestCoursesByFieldIdWhenInvalidToken(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, false) // No valid token

	fieldIdStr := "2"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusUnauthorized, res.StatusCode) // Expect 401 for invalid token
	mockCoursePageService.AssertNotCalled(t, "GetSuggestCourseByFieldId", fieldIdStr)
}

func TestGetCoursePageContentWhenRequestMethodIsInvalid(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	mockCoursePageId := "1"

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil) // Invalid method
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusMethodNotAllowed, res.StatusCode) // Expect 405 Method Not Allowed
	mockCoursePageService.AssertNotCalled(t, "GetCoursePageContent", mockCoursePageId)
}

func TestGetEnrollCourseByUserIdWhenUserIdIsInvalid(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	invalidUserId := "abc" // Invalid userId

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/enrollments/%s", invalidUserId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for invalid userId
	mockCoursePageService.AssertNotCalled(t, "GetEnrollCourseByUserId", invalidUserId)
}

func TestGetCoursePageContentWhenNoContent(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true)

	mockCoursePageId := "1"

	mockCoursePageService.On("GetCoursePageContent", mockCoursePageId).Return([]payload.CoursePageContent{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for empty content
	mockCoursePageService.AssertCalled(t, "GetCoursePageContent", mockCoursePageId)
}

func TestGetSuggestCoursesByFieldIdWhenFieldIdIsNotNumeric(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true)

	fieldIdStr := "non-numeric"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/suggest/%s", fieldIdStr), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404 for non-numeric fieldId
	mockCoursePageService.AssertNotCalled(t, "GetSuggestCourseByFieldId")
}

func TestGetSuggestCoursesByFieldIdWhenFieldIdIsEmpty(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true) // Inject token

	// Simulate a request with an empty fieldId
	req := httptest.NewRequest(http.MethodGet, "/courses/suggest/", nil) // No fieldId in the path
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	// Assert that the response returns 404 Not Found
	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode) // Expect 404
	mockCoursePageService.AssertNotCalled(t, "GetSuggestCourseByFieldId")
}

func TestGetCoursePageInfoWhenUnauthorized(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, false)

	mockCoursePageId := "1"

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/info", mockCoursePageId), nil)

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusUnauthorized, res.StatusCode) // Expect 401 for unauthorized request
	mockCoursePageService.AssertNotCalled(t, "GetCoursePageInfo", mockCoursePageId)
}

func TestGetCoursePageInfoWhenUnhandledError(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true)

	mockCoursePageId := "1"
	unexpectedError := fmt.Errorf("unexpected error occurred")

	mockCoursePageService.On("GetCoursePageInfo", mockCoursePageId).Return(nil, unexpectedError)

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

func TestGetCoursePageContentWhenEmptyResponse(t *testing.T) {
	is := assert.New(t)

	mockCoursePageService := new(mockServices.CoursePageServiceInterface)
	app := setupTestCoursePageController(mockCoursePageService, true)

	mockCoursePageId := "1"

	// Return an empty response
	mockCoursePageService.On("GetCoursePageContent", mockCoursePageId).Return([]payload.CoursePageContent{}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%s/content", mockCoursePageId), nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	res, err := app.Test(req)

	is.Nil(err)
	is.Equal(http.StatusNotFound, res.StatusCode)

	var responsePayload response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)
	is.Equal(fmt.Sprintf("no content found for course page ID %s", mockCoursePageId), responsePayload.Message)
}

