package controllers

import (
	"backend/internals/entities/payload"
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

type CourseControllerTestSuite struct {
	suite.Suite
}

func setupTestCourseController(mockCourseService *mockServices.CourseService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	courseController := NewCourseController(mockCourseService)

	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{"userId": float64(123)}
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/course/enrolled", courseController.GetEnrollCourseByUserId)
	return app
}

func (suite *CourseControllerTestSuite) TestGetEnrollCourseByUserIdWhenSuccess() {
	is := assert.New(suite.T())

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	course1 := "Course 1"
	course2 := "Course 2"
	course1Obj := &payload.Course{Name: &course1}
	course2Obj := &payload.Course{Name: &course2}
	mockEnrollments := []*payload.EnrollwithCourse{
		{Id: new(uint64), UserId: new(uint64), CourseId: new(uint64), CourseName: course1Obj, FieldName: new(string), FieldImageURL: new(string)},
		{Id: new(uint64), UserId: new(uint64), CourseId: new(uint64), CourseName: course2Obj, FieldName: new(string), FieldImageURL: new(string)},
	}

	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return(mockEnrollments, nil)

	req := httptest.NewRequest(http.MethodGet, "/course/enrolled", nil)
	res, err := app.Test(req)

	var responseBody response.InfoResponse[[]payload.EnrollwithCourse]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responseBody)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(len(mockEnrollments), len(responseBody.Data))
}

func (suite *CourseControllerTestSuite) TestGetEnrollCourseByUserIdWhenFailed() {
	is := assert.New(suite.T())

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return(nil, fmt.Errorf("failed to get enrollment info"))

	req := httptest.NewRequest(http.MethodGet, "/course/enrolled", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get enrollment info", errResponse.Message)
}

func TestCourseController(t *testing.T) {
	suite.Run(t, new(CourseControllerTestSuite))
}
