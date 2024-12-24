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

// Define Ptr functions to convert uint64 and string to pointers
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

	// Arrange
	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockUserID := uint(1)

	// Mock Course data with pointer types for fields
	mockCourse := models.Course{
		Id:   PtrUint64(10),                  // Use PtrUint64 to assign *uint64
		Name: PtrString("Test Course"),       // Use PtrString to assign *string
	}

	// Set up expectations for the mock repository method
	mockCourseRepo.EXPECT().
		GetCurrentCourse(mockUserID).
		Return(&mockCourse, nil)

	underTest := services.NewCourseService(mockCourseRepo)

	// Act
	course, err := underTest.GetCurrentCourse(mockUserID)

	// Assert
	is.NoError(err)
	is.Equal(mockCourse.Id, course.Id)
	is.Equal(mockCourse.Name, course.Name)
}

func (suite *CourseTestSuite) TestGetCurrentCourseFetchError() {
	is := assert.New(suite.T())

	// Arrange
	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockUserID := uint(1)

	// Set up expectation to mock an error when fetching current course
	mockCourseRepo.EXPECT().
		GetCurrentCourse(mockUserID).
		Return(nil, fmt.Errorf("failed to fetch current course"))

	underTest := services.NewCourseService(mockCourseRepo)

	// Act
	course, err := underTest.GetCurrentCourse(mockUserID)

	// Assert
	is.Error(err)
	is.Equal("failed to fetch current course", err.Error())
	is.Nil(course)
}

func (suite *CourseTestSuite) TestGetTotalStepsByCourseIdSuccess() {
    is := assert.New(suite.T())

    // Arrange
    mockCourseRepo := new(mockRepositories.CourseRepository)
    mockCourseID := uint(10)

    // Mock the return value for GetTotalStepsByCourseId
    expectedPayload := &payload.TotalStepsByCourseIdPayload{
        CourseId: mockCourseID,
        TotalSteps: 3,
    }
    mockCourseRepo.On("GetTotalStepsByCourseId", mockCourseID).Return(expectedPayload, nil)

    underTest := services.NewCourseService(mockCourseRepo)

    // Act
    actualPayload, err := underTest.GetTotalStepsByCourseId(mockCourseID)

    // Assert
    is.NoError(err)
    is.Equal(expectedPayload, actualPayload)  // Compare the entire struct instead of just the total steps
    mockCourseRepo.AssertExpectations(suite.T())  // Verify that the expectations were met
}


func (suite *CourseTestSuite) TestGetTotalStepsByCourseIdFetchError() {
	is := assert.New(suite.T())

	// Arrange
	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockCourseID := uint(10)

	// Set up expectation to mock an error when fetching total steps
	mockCourseRepo.EXPECT().
		GetTotalStepsByCourseId(mockCourseID).
		Return(0, fmt.Errorf("failed to fetch total steps"))

	underTest := services.NewCourseService(mockCourseRepo)

	// Act
	payload, err := underTest.GetTotalStepsByCourseId(mockCourseID)

	// Assert
	is.Error(err)
	is.Equal("failed to fetch total steps", err.Error())
	is.Nil(payload)
}

func TestCourseService(t *testing.T) {
	suite.Run(t, new(CourseTestSuite))
}
