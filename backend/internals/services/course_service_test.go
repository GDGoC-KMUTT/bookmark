package services_test

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func PtrUint64(val uint64) *uint64 {
	return &val
}

func PtrString(val string) *string {
	return &val
}

type CourseTestSuite struct {
	suite.Suite
}

func (suite *CourseTestSuite) TestGetCurrentCourseSuccess() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)
	mockUserID := uint(1)

	mockCourse := models.Course{
		Id:   PtrUint64(10),
		Name: PtrString("Test Course"),
	}

	mockCourseRepo.EXPECT().
		GetCurrentCourse(mockUserID).
		Return(&mockCourse, nil)

	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	course, err := underTest.GetCurrentCourse(mockUserID)

	is.NoError(err)
	is.Equal(mockCourse.Id, course.Id)
	is.Equal(mockCourse.Name, course.Name)
}

func (suite *CourseTestSuite) TestGetCurrentCourseFetchError() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	mockUserID := uint(1)

	mockCourseRepo.EXPECT().
		GetCurrentCourse(mockUserID).
		Return(nil, fmt.Errorf("failed to fetch current course"))

	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	course, err := underTest.GetCurrentCourse(mockUserID)

	is.Error(err)
	is.Equal("failed to fetch current course", err.Error())
	is.Nil(course)
}

func (suite *CourseTestSuite) TestGetTotalStepsByCourseIdSuccess() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	mockCourseID := uint(10)

	expectedPayload := &payload.TotalStepsByCourseIdPayload{
		CourseId:   mockCourseID,
		TotalSteps: 3,
	}

	mockCourseRepo.On("GetTotalStepsByCourseId", mockCourseID).Return(3, nil)

	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	actualPayload, err := underTest.GetTotalStepsByCourseId(mockCourseID)

	is.NoError(err)
	is.Equal(expectedPayload, actualPayload)

	mockCourseRepo.AssertExpectations(suite.T())
}

func (suite *CourseTestSuite) TestGetTotalStepsByCourseIdFetchError() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	mockCourseID := uint(10)

	mockCourseRepo.EXPECT().
		GetTotalStepsByCourseId(mockCourseID).
		Return(0, fmt.Errorf("failed to fetch total steps"))

	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	payload, err := underTest.GetTotalStepsByCourseId(mockCourseID)

	is.Error(err)
	is.Equal("failed to fetch total steps", err.Error())
	is.Nil(payload)
}

func (suite *CourseTestSuite) TestGetCourseByFieldIdWhenSuccess() {
	is := assert.New(suite.T())

	// Arrange
	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	// Mock data
	mockCourseId := utils.Ptr(uint64(1))
	mockCourseName := utils.Ptr("testname")
	mockCourseFieldId := utils.Ptr(uint(1))
	mockFieldId := utils.Ptr(uint64(1))
	mockFieldName := utils.Ptr("testname")
	mockFieldImageUrl := utils.Ptr("testimageurl")

	mockCourseRepo.EXPECT().FindCourseByFieldId(mockCourseFieldId).Return([]models.Course{{
		Id:      mockCourseId,
		Name:    mockCourseName,
		FieldId: mockFieldId,
	}}, &models.FieldType{
		Id:       mockFieldId,
		Name:     mockFieldName,
		ImageUrl: mockFieldImageUrl,
	}, nil)

	// Test
	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)
	courses, err := underTest.GetCourseByFieldId(mockCourseFieldId)
	is.Equal(courses[0].Id, mockCourseId)
	is.Equal(courses[0].Name, mockCourseName)
	is.Equal(courses[0].FieldId, mockFieldId)
	is.Equal(courses[0].FieldName, mockFieldName)
	is.Equal(courses[0].FieldImageUrl, mockFieldImageUrl)
	is.NoError(err)

}
func (suite *CourseTestSuite) TestGetCourseByFieldIdWhenFailed() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	mockFieldId := utils.Ptr(uint(1))

	mockCourseRepo.On("FindCourseByFieldId", mockFieldId).Return(nil, nil, fmt.Errorf("courses not found"))
	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	courses, err := underTest.GetCourseByFieldId(mockFieldId)

	is.Nil(courses)
	is.NotNil(err)
}

func (suite *CourseTestSuite) TestGetAllFieldTypesWhenSuccess() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	mockId := utils.Ptr(uint64(1))
	mockName := utils.Ptr("testname")
	mockImageUrl := utils.Ptr("testimageurl")

	mockFieldTypeRepo.On("FindAllFieldTypes").Return([]models.FieldType{{Id: mockId,
		Name:     mockName,
		ImageUrl: mockImageUrl},
	}, nil)

	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	fieldTypes, err := underTest.GetAllFieldTypes()

	is.NoError(err)
	is.Len(fieldTypes, 1)
	is.Equal(fieldTypes[0].Id, mockId)
	is.Equal(fieldTypes[0].Name, mockName)
	is.Equal(fieldTypes[0].ImageUrl, mockImageUrl)
}

func (suite *CourseTestSuite) TestGetAllFieldTypesWhenFailed() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockFieldTypeRepo := new(mockRepositories.FieldTypeRepository)

	mockFieldTypeRepo.On("FindAllFieldTypes").Return(nil, fmt.Errorf("fieldTypes not found"))
	underTest := services.NewCourseService(mockCourseRepo, mockFieldTypeRepo)

	fieldTypes, err := underTest.GetAllFieldTypes()

	is.Nil(fieldTypes)
	is.NotNil(err)

}

func TestCourseService(t *testing.T) {
	suite.Run(t, new(CourseTestSuite))
}
