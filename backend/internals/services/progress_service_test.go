package services_test

import (
	"backend/internals/db/models"
	"backend/internals/services"
	mockRepositories "backend/mocks/repositories"
	"backend/internals/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProgressTestSuite struct {
	suite.Suite
}

func (suite *ProgressTestSuite) TestGetCompletionPercentageSuccess() {
	is := assert.New(suite.T())

	mockUserRepo := new(mockRepositories.UserRepository)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	mockUserID := uint(1)
	mockCourseID := uint(10)

	mockSteps := []models.Step{
		{Id: utils.Ptr(uint64(1))},
		{Id: utils.Ptr(uint64(2))},
		{Id: utils.Ptr(uint64(3))},
	}

	mockUserActivities := []models.UserActivity{
		{StepId: utils.Ptr(uint64(1))},
		{StepId: utils.Ptr(uint64(3))},
	}

	mockCourseRepo.EXPECT().
		GetAllCourseSteps(mockCourseID).
		Return(mockSteps, nil)

	mockUserRepo.EXPECT().
		GetUserCompletedSteps(mockUserID).
		Return(mockUserActivities, nil)

	underTest := services.NewProgressService(mockUserRepo, mockCourseRepo)

	percentage, err := underTest.GetCompletionPercentage(mockUserID, mockCourseID)

	is.NoError(err)
	is.InDelta(66.67, percentage, 0.01)
}

func (suite *ProgressTestSuite) TestGetCompletionPercentageNoStepsFound() {
	is := assert.New(suite.T())

	mockUserRepo := new(mockRepositories.UserRepository)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	mockUserID := uint(1)
	mockCourseID := uint(10)

	mockCourseRepo.EXPECT().
		GetAllCourseSteps(mockCourseID).
		Return([]models.Step{}, nil)

	mockUserRepo.EXPECT().
		GetUserCompletedSteps(mockUserID).
		Return([]models.UserActivity{}, nil)

	underTest := services.NewProgressService(mockUserRepo, mockCourseRepo)

	percentage, err := underTest.GetCompletionPercentage(mockUserID, mockCourseID)

	is.Error(err)
	is.Equal("no steps found for course ID 10", err.Error())
	is.Equal(0.0, percentage)
}

func (suite *ProgressTestSuite) TestGetCompletionPercentageFetchError() {
	is := assert.New(suite.T())

	mockUserRepo := new(mockRepositories.UserRepository)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	mockUserID := uint(1)
	mockCourseID := uint(10)

	mockCourseRepo.EXPECT().
		GetAllCourseSteps(mockCourseID).
		Return(nil, fmt.Errorf("failed to fetch course steps"))

	underTest := services.NewProgressService(mockUserRepo, mockCourseRepo)

	percentage, err := underTest.GetCompletionPercentage(mockUserID, mockCourseID)

	is.Error(err)
	is.Equal("failed to fetch course steps", err.Error())
	is.Equal(0.0, percentage)
}

func TestProgressService(t *testing.T) {
	suite.Run(t, new(ProgressTestSuite))
}
