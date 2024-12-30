package controllers_test

import (
	"backend/internals/controllers"
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

type EnrollControllerTestSuite struct {
	suite.Suite
}

func setupTestEnrollController(mockService *mockServices.EnrollServices) *fiber.App {
	app := fiber.New()
	controller := controllers.NewEnrollController(mockService)
	app.Post("/enroll/:userId/:courseId", controller.EnrollInCourse)
	return app
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenSuccess() {
	is := assert.New(suite.T())

	mockService := new(mockServices.EnrollServices)
	app := setupTestEnrollController(mockService)

	mockUserId := uint64(123)
	mockCourseId := uint64(456)

	mockService.EXPECT().EnrollUser(mockUserId, mockCourseId).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/enroll/123/456", nil)
	res, err := app.Test(req)

	var response map[string]string
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal("user enrolled successfully", response["message"])
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenInvalidUserId() {
	is := assert.New(suite.T())

	mockService := new(mockServices.EnrollServices)
	app := setupTestEnrollController(mockService)

	req := httptest.NewRequest(http.MethodPost, "/enroll/abc/456", nil)
	res, err := app.Test(req)

	var response map[string]string
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("invalid user ID", response["error"])
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenUserAlreadyEnrolled() {
	is := assert.New(suite.T())

	mockService := new(mockServices.EnrollServices)
	app := setupTestEnrollController(mockService)

	mockUserId := uint64(123)
	mockCourseId := uint64(456)

	mockService.EXPECT().EnrollUser(mockUserId, mockCourseId).Return(fmt.Errorf("user is already enrolled in this course"))

	req := httptest.NewRequest(http.MethodPost, "/enroll/123/456", nil)
	res, err := app.Test(req)

	var response map[string]string
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)

	is.Nil(err)
	is.Equal(http.StatusConflict, res.StatusCode)
	is.Equal("user already enrolled", response["error"])
}

func (suite *EnrollControllerTestSuite) TestEnrollInCourseWhenInternalError() {
	is := assert.New(suite.T())

	mockService := new(mockServices.EnrollServices)
	app := setupTestEnrollController(mockService)

	mockUserId := uint64(123)
	mockCourseId := uint64(456)

	mockService.EXPECT().EnrollUser(mockUserId, mockCourseId).Return(fmt.Errorf("some internal error"))

	req := httptest.NewRequest(http.MethodPost, "/enroll/123/456", nil)
	res, err := app.Test(req)

	var response map[string]string
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("internal server error", response["error"])
}

func TestEnrollController(t *testing.T) {
	suite.Run(t, new(EnrollControllerTestSuite))
}
