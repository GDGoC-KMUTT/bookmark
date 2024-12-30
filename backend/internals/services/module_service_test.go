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

type ModuleServiceTestSuite struct {
	suite.Suite
}

func (suite *ModuleServiceTestSuite) TestGetModuleInfoWhenSuccess() {
	is := assert.New(suite.T())

	// Arrange
	mockModuleRepo := new(mockRepositories.ModuleRepo)
	mockModuleID := "123"
	mockModuleEntity := &models.Module{
		Id:          utils.Ptr(uint64(123)),
		Title:       utils.Ptr("Test Module"),
		Description: utils.Ptr("This is a description."),
		ImageUrl:    utils.Ptr("http://example.com/image.png"),
	}

	mockModuleRepo.On("FindModuleInfoByModuleID", mockModuleID).Return(mockModuleEntity, nil)

	// Test
	underTest := services.NewModuleService(mockModuleRepo)
	moduleInfo, err := underTest.GetModuleInfo(mockModuleID)

	// Assert
	is.NoError(err)
	is.NotNil(moduleInfo)
	is.Equal(uint64(123), moduleInfo.Id)
	is.Equal("Test Module", moduleInfo.Title)
	is.Equal("This is a description.", *moduleInfo.Description)
	is.Equal("http://example.com/image.png", *moduleInfo.ImageUrl)
}

func (suite *ModuleServiceTestSuite) TestGetModuleInfoWhenRepoFails() {
	is := assert.New(suite.T())

	// Arrange
	mockModuleRepo := new(mockRepositories.ModuleRepo)
	mockModuleID := "123"

	mockModuleRepo.On("FindModuleInfoByModuleID", mockModuleID).Return(nil, fmt.Errorf("repository error"))

	// Test
	underTest := services.NewModuleService(mockModuleRepo)
	moduleInfo, err := underTest.GetModuleInfo(mockModuleID)

	// Assert
	is.Nil(moduleInfo)
	is.NotNil(err)
	is.Equal("repository error", err.Error())
}

func (suite *ModuleServiceTestSuite) TestGetModuleInfoWhenRepoReturnsNilEntity() {
	is := assert.New(suite.T())

	// Arrange
	mockModuleRepo := new(mockRepositories.ModuleRepo)
	mockModuleID := "123"

	mockModuleRepo.On("FindModuleInfoByModuleID", mockModuleID).Return(nil, nil)

	// Test
	underTest := services.NewModuleService(mockModuleRepo)
	moduleInfo, err := underTest.GetModuleInfo(mockModuleID)

	// Assert
	is.Nil(moduleInfo)
	is.Nil(err) // Ensure no error is returned when the entity is nil
}

func TestModuleService(t *testing.T) {
	suite.Run(t, new(ModuleServiceTestSuite))
}
