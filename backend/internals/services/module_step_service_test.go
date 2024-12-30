package services_test

import (
	"backend/internals/db/models"
	"backend/internals/utils"
	"backend/internals/services"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ModuleStepServiceTestSuite struct {
	suite.Suite
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsWhenSuccess() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.StepRepo)
	mockModuleID := "123"
	mockSteps := []models.Step{
		{
			Id:    utils.Ptr(uint64(1)),
			Title: utils.Ptr("Step 1"),
			Check: utils.Ptr("true"),
		},
		{
			Id:    utils.Ptr(uint64(2)),
			Title: utils.Ptr("Step 2"),
			Check: utils.Ptr("false"),
		},
	}

	mockRepo.On("FindStepsByModuleID", mockModuleID).Return(mockSteps, nil)

	// Test
	underTest := services.NewModuleStepService(mockRepo)
	steps, err := underTest.GetModuleSteps(mockModuleID)

	// Assert
	is.NoError(err)
	is.NotNil(steps)
	is.Len(steps, 2)
	is.Equal(uint64(1), steps[0].Id)
	is.Equal("Step 1", steps[0].Title)
	is.Equal("true", steps[0].Check)
	is.Equal(uint64(2), steps[1].Id)
	is.Equal("Step 2", steps[1].Title)
	is.Equal("false", steps[1].Check)
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsWhenRepoFails() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.StepRepo)
	mockModuleID := "123"

	mockRepo.On("FindStepsByModuleID", mockModuleID).Return(nil, fmt.Errorf("repository error"))

	// Test
	underTest := services.NewModuleStepService(mockRepo)
	steps, err := underTest.GetModuleSteps(mockModuleID)

	// Assert
	is.Nil(steps)
	is.NotNil(err)
	is.Equal("repository error", err.Error())
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsWhenRepoReturnsNoSteps() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.StepRepo)
	mockModuleID := "123"

	mockRepo.On("FindStepsByModuleID", mockModuleID).Return([]models.Step{}, nil)

	// Test
	underTest := services.NewModuleStepService(mockRepo)
	steps, err := underTest.GetModuleSteps(mockModuleID)

	// Assert
	is.NoError(err)
	is.Len(steps, 0)
}

func TestModuleStepService(t *testing.T) {
	suite.Run(t, new(ModuleStepServiceTestSuite))
}
