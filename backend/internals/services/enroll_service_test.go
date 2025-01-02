package services_test

import (
	"backend/internals/services"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnrollServiceTestSuite struct {
	suite.Suite
}

func TestEnrollService(t *testing.T) {
	suite.Run(t, new(EnrollServiceTestSuite))
}

func (suite *EnrollServiceTestSuite) TestEnrollUserWhenSuccess() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.EnrollRepository)
	mockUserId := uint(123)
	mockCourseId := uint64(456)

	mockRepo.On("EnrollUser", mockUserId, mockCourseId).Return(nil)

	// Test
	underTest := services.NewEnrollService(mockRepo)
	err := underTest.EnrollUser(mockUserId, mockCourseId)

	// Assert
	is.NoError(err)
	//mockRepo.AssertCalled(t, "EnrollUser", mockUserId, mockCourseId)
}

func (suite *EnrollServiceTestSuite) TestEnrollUserWhenRepoFails() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.EnrollRepository)
	mockUserId := uint(123)
	mockCourseId := uint64(456)

	mockRepo.On("EnrollUser", mockUserId, mockCourseId).Return(fmt.Errorf("repository error"))

	// Test
	underTest := services.NewEnrollService(mockRepo)
	err := underTest.EnrollUser(mockUserId, mockCourseId)

	// Assert
	is.NotNil(err)
	is.Equal("repository error", err.Error())
	//mockRepo.AssertCalled(t, "EnrollUser", mockUserId, mockCourseId)
}
