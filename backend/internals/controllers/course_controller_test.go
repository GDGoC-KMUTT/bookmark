package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	mockServices "backend/mocks/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupTestCourseController(courseSvc services.CourseService) *fiber.App {
	app := fiber.New()

	controller := NewCourseController(courseSvc)

	app.Use(func(c *fiber.Ctx) error {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["userId"] = float64(123)
		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/courses/current", controller.GetCurrentCourse)
	app.Get("/courses/:courseId/total-steps", controller.GetTotalStepsByCourseId)

	return app
}

func TestGetCurrentCourseWhenSuccess(t *testing.T) {
    is := assert.New(t)

    mockCourseService := new(mockServices.CourseService)

    app := setupTestCourseController(mockCourseService)

    mockUserId := uint(123)
    mockCourseId := uint64(1)
    mockCourseName := "Test Course"
    
    expectedCourse := payload.Course{
        Id:   &mockCourseId,
        Name: &mockCourseName,
    }

    mockCourseService.EXPECT().GetCurrentCourse(mockUserId).Return(&expectedCourse, nil)

    req := httptest.NewRequest(http.MethodGet, "/courses/current", nil)
    req.Header.Set("Authorization", "Bearer mockToken")

    res, err := app.Test(req)

    var responsePayload response.InfoResponse[payload.Course]
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &responsePayload)

    is.Nil(err)
    is.Equal(http.StatusOK, res.StatusCode)

    is.Equal(*expectedCourse.Id, *responsePayload.Data.Id)
    is.Equal(*expectedCourse.Name, *responsePayload.Data.Name)
}

func TestGetCurrentCourseWhenFailedToFetchCurrentCourse(t *testing.T) {
	is := assert.New(t)

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockUserId := uint(123)
	mockCourseService.EXPECT().GetCurrentCourse(mockUserId).Return(nil, fmt.Errorf("failed to fetch current course"))

	req := httptest.NewRequest(http.MethodGet, "/courses/current", nil)
	req.Header.Set("Authorization", "Bearer mockToken")

	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	// is.Equal("failed to fetch current course", errResponse.Message)
}

func TestGetTotalStepsByCourseIdWhenSuccess(t *testing.T) {
    is := assert.New(t)

    mockCourseService := new(mockServices.CourseService)

    app := setupTestCourseController(mockCourseService)

    mockCourseId := uint(1)
    expectedTotalSteps := payload.TotalStepsByCourseIdPayload{
        TotalSteps: 10,
    }

    // Expect the incremented courseId, which will be mockCourseId + 1
    mockCourseService.EXPECT().GetTotalStepsByCourseId(mockCourseId).Return(&expectedTotalSteps, nil)

    req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%d/total-steps", mockCourseId), nil)
    res, err := app.Test(req)

    var responsePayload response.InfoResponse[payload.TotalStepsByCourseIdPayload]
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &responsePayload)

    is.Nil(err)
    is.Equal(http.StatusOK, res.StatusCode)
    is.Equal(expectedTotalSteps.TotalSteps, responsePayload.Data.TotalSteps)
}

func TestGetTotalStepsByCourseIdWhenFailedToFetchTotalSteps(t *testing.T) {
	is := assert.New(t)

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockCourseId := uint(1)

	mockCourseService.EXPECT().GetTotalStepsByCourseId(mockCourseId - 1).Return(&payload.TotalStepsByCourseIdPayload{}, fmt.Errorf("failed to fetch total steps"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%d/total-steps", mockCourseId), nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	// is.Equal("failed to fetch total steps", errResponse.Message)
}
