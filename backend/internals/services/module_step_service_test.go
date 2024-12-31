package services_test

import (
	"backend/internals/db/models"
	mockRepositories "backend/mocks/repositories"
	"backend/internals/services"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ModuleStepServiceTestSuite struct {
	suite.Suite
	mockStepRepo     *mockRepositories.StepRepository
	mockUserEvalRepo *mockRepositories.UserEvaluateRepository
	service          services.ModuleStepServices
}

func (suite *ModuleStepServiceTestSuite) SetupTest() {
	suite.mockStepRepo = mockRepositories.NewStepRepository(suite.T())
	suite.mockUserEvalRepo = mockRepositories.NewUserEvaluateRepository(suite.T())
	suite.service = services.NewModuleStepService(suite.mockStepRepo, suite.mockUserEvalRepo)
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsSuccess() {
	moduleID := "module123"
	steps := []*models.Step{
		{Id: new(uint64), Title: new(string)},
	}
	*steps[0].Id = 1
	*steps[0].Title = "Step 1"

	suite.mockStepRepo.EXPECT().FindStepsByModuleID(&moduleID).Return(steps, nil)
	suite.mockUserEvalRepo.EXPECT().FindStepEvaluateIDsByStepID(uint64(1)).Return([]uint64{1, 2}, nil)
	suite.mockUserEvalRepo.EXPECT().FindUserPassedEvaluateIDs(uint(1), uint64(1)).Return([]uint64{1, 2}, nil)

	result, err := suite.service.GetModuleSteps(1, moduleID)
	suite.NoError(err)
	suite.Len(result, 1)
	suite.True(result[0].Check)
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsNoStepsFound() {
	moduleID := "module123"
	suite.mockStepRepo.EXPECT().FindStepsByModuleID(&moduleID).Return(nil, nil)

	result, err := suite.service.GetModuleSteps(1, moduleID)
	suite.Error(err)
	suite.Nil(result)
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsEvaluationMismatch() {
	moduleID := "module123"
	steps := []*models.Step{
		{Id: new(uint64), Title: new(string)},
	}
	*steps[0].Id = 1
	*steps[0].Title = "Step 1"

	suite.mockStepRepo.EXPECT().FindStepsByModuleID(&moduleID).Return(steps, nil)
	suite.mockUserEvalRepo.EXPECT().FindStepEvaluateIDsByStepID(uint64(1)).Return([]uint64{1, 2}, nil)
	suite.mockUserEvalRepo.EXPECT().FindUserPassedEvaluateIDs(uint(1), uint64(1)).Return([]uint64{1}, nil)

	result, err := suite.service.GetModuleSteps(1, moduleID)
	suite.NoError(err)
	suite.Len(result, 1)
	suite.False(result[0].Check)
}

func (suite *ModuleStepServiceTestSuite) TestGetModuleStepsRepoError() {
	moduleID := "module123"
	suite.mockStepRepo.EXPECT().FindStepsByModuleID(&moduleID).Return(nil, errors.New("repository error"))

	result, err := suite.service.GetModuleSteps(1, moduleID)
	suite.Error(err)
	suite.Nil(result)
}

func TestModuleStepService(t *testing.T) {
	suite.Run(t, new(ModuleStepServiceTestSuite))
}
