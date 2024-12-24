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
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)

	// Mock data
	mockUserID := uint(1) // Use uint instead of uint64 to match the method signature
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
		Gems:        utils.Ptr(mockGems),
	}

	mockUserPass := &models.UserPass{
		Id:     utils.Ptr(uint64(1)),
		UserId: utils.Ptr(uint64(mockUserID)), // Ensure the type matches `uint64` here for compatibility
		StepId: utils.Ptr(mockStepID),
		Step:   mockStep,
		Type:   utils.Ptr("step"),
	}

	// Mock repository behavior
	mockUserRepo.EXPECT().
		GetTotalGemsByUserID(mockUserID).
		Return(mockUserPass.Step.Gems, nil) // Use the gems from the step in mockUserPass

	// Test
	underTest := services.NewProfileService(mockUserRepo)

	// Act
	totalGems, err := underTest.GetTotalGems(mockUserID)

	// Assert
	is.NoError(err)
	is.Equal(mockUserID, totalGems.UserID)
	is.Equal(mockGems, totalGems.Total)
}

func (suite *ProfileTestSuit) TestGetTotalGemsWhenFailed() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)

	// Mock data
	mockUserID := uint(1)

	// mock user repo
	mockUserRepo.EXPECT().GetTotalGemsByUserID(mockUserID).Return(0, fmt.Errorf("gems data not found"))

	// Test
	underTest := services.NewProfileService(mockUserRepo)

	// Test Failure
	gemTotal, err := underTest.GetTotalGems(mockUserID)

	is.Nil(gemTotal)
	is.NotNil(err)
}

func TestProfileService(t *testing.T) {
	suite.Run(t, new(ProfileTestSuit))
}
