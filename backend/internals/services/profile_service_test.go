package services_test

import (
	"backend/internals/db/models"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type ProfileTestSuit struct {
	suite.Suite
}

func (suite *ProfileTestSuit) TestGetUserInfoWhenSuccess() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)

	// Mock data
	mockUserId := utils.Ptr[uint64](1)
	mockFirstName := utils.Ptr("testfn")
	mockLastName := utils.Ptr("testln")
	mockEmail := utils.Ptr("test@gmail.com")
	mockPhotoUrl := utils.Ptr("url")

	// mock user repo
	mockUserRepo.EXPECT().FindUserByID(mock.Anything).Return(&models.User{
		Id:        mockUserId,
		Firstname: mockFirstName,
		Lastname:  mockLastName,
		Email:     mockEmail,
		PhotoUrl:  mockPhotoUrl,
	}, nil)

	// Test
	underTest := services.NewProfileService(mockUserRepo)

	// Test Success
	userInfo, err := underTest.GetUserInfo(utils.Ptr(strconv.Itoa(int(*mockUserId))))

	is.Equal(*mockUserId, *userInfo.Id)
	is.Equal(*mockFirstName, *userInfo.Firstname)
	is.NoError(err)
}

func (suite *ProfileTestSuit) TestGetUserInfoWhenFailed() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)

	// Mock data
	mockUserId := utils.Ptr[uint64](1)

	// mock user repo
	mockUserRepo.EXPECT().FindUserByID(mock.Anything).Return(nil, fmt.Errorf("user not found"))

	// Test
	underTest := services.NewProfileService(mockUserRepo)

	// Test Success
	userInfo, err := underTest.GetUserInfo(utils.Ptr(strconv.Itoa(int(*mockUserId))))

	is.Nil(userInfo)
	is.NotNil(err)
}

func (suite *ProfileTestSuit) TestGetTotalGemsWhenSuccess() {
	is := assert.New(suite.T())
	mockUserRepo := new(mockRepositories.UserRepository)

	mockUserID := uint(1)
	mockModuleID := uint64(101)
	mockStepID := uint64(201)
	mockGems := int64(50)

	mockModule := &models.Module{
		Id:          utils.Ptr(mockModuleID),
		Title:       utils.Ptr("Test Module"),
		Description: utils.Ptr("This is a test module"),
	}

	mockStep := &models.Step{
		Id:          utils.Ptr(mockStepID),
		ModuleId:    utils.Ptr(mockModuleID),
		Module:      mockModule,
		Title:       utils.Ptr("Test Step"),
		Description: utils.Ptr("This is a test step"),
	}

	mockStepEvaluate := &models.StepEvaluate{
		Id:     utils.Ptr(uint64(1)),
		StepId: utils.Ptr(mockStepID),
		Step:   mockStep,
		Gem:    utils.Ptr(int(mockGems)),
	}

	var totalGems uint64
	if mockStepEvaluate.Gem != nil {
		totalGems = uint64(*mockStepEvaluate.Gem)
	}

	mockUserRepo.EXPECT().
		GetTotalGemsByUserID(mockUserID).
		Return(totalGems, nil)

	underTest := services.NewProfileService(mockUserRepo)

	result, err := underTest.GetTotalGems(mockUserID)

	is.NoError(err)
	is.Equal(mockUserID, result.UserID)
	is.Equal(totalGems, result.Total)
}

func (suite *ProfileTestSuit) TestGetTotalGemsWhenFailed() {
	is := assert.New(suite.T())
	mockUserRepo := new(mockRepositories.UserRepository)

	mockUserID := uint(1)

	mockUserRepo.EXPECT().GetTotalGemsByUserID(mockUserID).Return(0, fmt.Errorf("gems data not found"))

	underTest := services.NewProfileService(mockUserRepo)

	gemTotal, err := underTest.GetTotalGems(mockUserID)

	is.Nil(gemTotal)
	is.NotNil(err)
}

func TestProfileService(t *testing.T) {
	suite.Run(t, new(ProfileTestSuit))
}
