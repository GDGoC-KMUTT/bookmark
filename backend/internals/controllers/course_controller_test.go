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

type CourseControllerTestSuit struct {
	suite.Suite
}

func setupTestCourseController(courseSvc services.CourseService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

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
	app.Get("/courses/field-types", controller.GetAllFieldTypes)
	app.Get("/courses/field/:fieldId", controller.GetCoursesByFieldId)
	app.Get("/courses/enrolled", controller.GetEnrollCourseByUserId)
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

	mockCourseService.EXPECT().GetTotalStepsByCourseId(mockCourseId).Return(&payload.TotalStepsByCourseIdPayload{}, fmt.Errorf("failed to fetch total steps"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/%d/total-steps", mockCourseId), nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	// is.Equal("failed to fetch total steps", errResponse.Message)
}

// func TestGetEnrollCourseByUserIdWhenNoEnrollments(t *testing.T) {
// 	is := assert.New(t)

// 	mockCourseService := new(mockServices.CourseService)
// 	app := setupTestCourseController(mockCourseService)

// 	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return([]*payload.EnrollwithCourse{}, nil)

// 	req := httptest.NewRequest(http.MethodGet, "/courses/enrolled", nil)
// 	res, err := app.Test(req)

// 	var responsePayload response.InfoResponse[[]*payload.EnrollwithCourse]
// 	body, _ := io.ReadAll(res.Body)
// 	json.Unmarshal(body, &responsePayload)

// 	is.Nil(err)
// 	is.Equal(http.StatusOK, res.StatusCode)
// 	is.Empty(responsePayload.Data)
// }

// func TestGetEnrollCourseByUserIdWhenEnrollmentsFound(t *testing.T) {
// 	is := assert.New(t)

// 	mockCourseService := new(mockServices.CourseService)
// 	app := setupTestCourseController(mockCourseService)

// 	mockEnrollments := []*payload.EnrollwithCourse{
// 	}

// 	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return(mockEnrollments, nil)

// 	req := httptest.NewRequest(http.MethodGet, "/courses/enrolled", nil)
// 	res, err := app.Test(req)

// 	var responsePayload response.InfoResponse[[]*payload.EnrollwithCourse]
// 	body, _ := io.ReadAll(res.Body)
// 	json.Unmarshal(body, &responsePayload)

// 	is.Nil(err)
// 	is.Equal(http.StatusOK, res.StatusCode)
// 	is.Equal(len(mockEnrollments), len(responsePayload.Data))
// }

// func TestGetEnrollCourseByUserIdWhenServiceError(t *testing.T) {
// 	is := assert.New(t)

// 	mockCourseService := new(mockServices.CourseService)
// 	app := setupTestCourseController(mockCourseService)

// 	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return(nil, fmt.Errorf("failed to fetch enrollments"))

// 	req := httptest.NewRequest(http.MethodGet, "/courses/enrolled", nil)

// }

func (suite *CourseControllerTestSuit) TestGetCourseByFieldIdWhenSuccess() {
	is := assert.New(suite.T())

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockId := uint64(1)
	mockName := "testname"
	mockFieldId := uint(1)
	mockCourseFieldId := uint64(1)
	mockFieldName := "testfieldname"
	mockFieldImageUrl := "testfieldimageurl"

	mockCourseService.EXPECT().GetCoursesByFieldId(mockFieldId).Return([]payload.CourseWithFieldType{{
		Id:            utils.Ptr(mockId),
		Name:          utils.Ptr(mockName),
		FieldId:       utils.Ptr(mockCourseFieldId),
		FieldName:     utils.Ptr(mockFieldName),
		FieldImageUrl: utils.Ptr(mockFieldImageUrl),
	},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/field/%d", mockFieldId), nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[[]payload.CourseWithFieldType])
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(mockId, *r.Data[0].Id)
	is.Equal(mockName, *r.Data[0].Name)
	is.Equal(mockCourseFieldId, *r.Data[0].FieldId)
	is.Equal(mockFieldName, *r.Data[0].FieldName)
	is.Equal(mockFieldImageUrl, *r.Data[0].FieldImageUrl)
	is.Equal(http.StatusOK, res.StatusCode)

}

func (suite *CourseControllerTestSuit) TestGetCourseByFieldIdWhenFailedToFetchCourseByFieldId() {
	is := assert.New(suite.T())

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockFieldId := uint(1)
	mockCourseService.EXPECT().GetCoursesByFieldId(mockFieldId).Return(nil, fmt.Errorf("failed to get course by fieldId"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/courses/field/%d", mockFieldId), nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get course by fieldId", errResponse.Message)
}

func (suite *CourseControllerTestSuit) TestGetAllFieldTypesWhenSuccess() {
	is := assert.New(suite.T())

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockId := uint64(1)
	mockName := "testname"
	mockImageUrl := "testimageurl"
	mockCourseService.EXPECT().GetAllFieldTypes().Return([]payload.FieldType{
		{
			Id:       utils.Ptr(mockId),
			Name:     utils.Ptr(mockName),
			ImageUrl: utils.Ptr(mockImageUrl),
		},
	}, nil)
	req := httptest.NewRequest(http.MethodGet, "/courses/field-types", nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[[]payload.FieldType])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(mockId, *r.Data[0].Id)
	is.Equal(mockName, *r.Data[0].Name)
	is.Equal(mockImageUrl, *r.Data[0].ImageUrl)
	is.Equal(http.StatusOK, res.StatusCode)

}

func (suite *CourseControllerTestSuit) TestGetAllFieldTypesWheFailedToGetAllFieldTypes() {
	is := assert.New(suite.T())

	mockCourseService := new(mockServices.CourseService)

	app := setupTestCourseController(mockCourseService)

	mockCourseService.EXPECT().GetAllFieldTypes().Return(nil, fmt.Errorf("failed to get all field types"))

	req := httptest.NewRequest(http.MethodGet, "/courses/field-types", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get all field types", errResponse.Message)
}

func TestGetEnrollCourseByUserIdWhenNoEnrollments(t *testing.T) {
	is := assert.New(t)

	mockCourseService := new(mockServices.CourseService)
	app := setupTestCourseController(mockCourseService)

	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return([]*payload.EnrollwithCourse{}, nil)

	req := httptest.NewRequest(http.MethodGet, "/courses/enrolled", nil)
	res, err := app.Test(req)

	var responsePayload response.InfoResponse[[]*payload.EnrollwithCourse]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Empty(responsePayload.Data)
}

func TestGetEnrollCourseByUserIdWhenEnrollmentsFound(t *testing.T) {
	is := assert.New(t)

	mockCourseService := new(mockServices.CourseService)
	app := setupTestCourseController(mockCourseService)

	mockEnrollments := []*payload.EnrollwithCourse{}

	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return(mockEnrollments, nil)

	req := httptest.NewRequest(http.MethodGet, "/courses/enrolled", nil)
	res, err := app.Test(req)

	var responsePayload response.InfoResponse[[]*payload.EnrollwithCourse]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(len(mockEnrollments), len(responsePayload.Data))
}

func TestGetEnrollCourseByUserIdWhenServiceError(t *testing.T) {
	is := assert.New(t)

	mockCourseService := new(mockServices.CourseService)
	app := setupTestCourseController(mockCourseService)

	mockCourseService.EXPECT().GetEnrollCourseByUserId(mock.Anything).Return(nil, fmt.Errorf("failed to fetch enrollments"))

	req := httptest.NewRequest(http.MethodGet, "/courses/enrolled", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	// is.Equal("Failed to fetch enrollments", errResponse.Message)


}
