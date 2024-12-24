package services_test

import (
	"backend/internals/db/models"
	"backend/internals/services"
	"backend/internals/entities/payload"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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
	mockUserID := uint(1)

	mockCourse := models.Course{
		Id:   PtrUint64(10),
		Name: PtrString("Test Course"),
	}

	mockCourseRepo.EXPECT().
		GetCurrentCourse(mockUserID).
		Return(&mockCourse, nil)

	underTest := services.NewCourseService(mockCourseRepo)

	course, err := underTest.GetCurrentCourse(mockUserID)

	is.NoError(err)
	is.Equal(mockCourse.Id, course.Id)
	is.Equal(mockCourse.Name, course.Name)
}

func (suite *CourseTestSuite) TestGetCurrentCourseFetchError() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockUserID := uint(1)

	mockCourseRepo.EXPECT().
		GetCurrentCourse(mockUserID).
		Return(nil, fmt.Errorf("failed to fetch current course"))

	underTest := services.NewCourseService(mockCourseRepo)

	course, err := underTest.GetCurrentCourse(mockUserID)

	is.Error(err)
	is.Equal("failed to fetch current course", err.Error())
	is.Nil(course)
}

func (suite *CourseTestSuite) TestGetTotalStepsByCourseIdSuccess() {
    is := assert.New(suite.T())

    mockCourseRepo := new(mockRepositories.CourseRepository)
    mockCourseID := uint(10)

    expectedPayload := &payload.TotalStepsByCourseIdPayload{
        CourseId: mockCourseID,
        TotalSteps: 3,
    }

    mockCourseRepo.On("GetTotalStepsByCourseId", mockCourseID).Return(3, nil)

    underTest := services.NewCourseService(mockCourseRepo)

    actualPayload, err := underTest.GetTotalStepsByCourseId(mockCourseID)

    is.NoError(err)
    is.Equal(expectedPayload, actualPayload)

    mockCourseRepo.AssertExpectations(suite.T())
}

func (suite *CourseTestSuite) TestGetTotalStepsByCourseIdFetchError() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockCourseID := uint(10)

	mockCourseRepo.EXPECT().
		GetTotalStepsByCourseId(mockCourseID).
		Return(0, fmt.Errorf("failed to fetch total steps"))

	underTest := services.NewCourseService(mockCourseRepo)

	payload, err := underTest.GetTotalStepsByCourseId(mockCourseID)

	is.Error(err)
	is.Equal("failed to fetch total steps", err.Error())
	is.Nil(payload)
}

func TestCourseService(t *testing.T) {
	suite.Run(t, new(CourseTestSuite))
}
