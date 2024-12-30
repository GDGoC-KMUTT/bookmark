package services_test

import (
	"backend/internals/services"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnrollServiceTestSuite struct {
	suite.Suite
}

func (suite *EnrollServiceTestSuite) TestEnrollUserWhenSuccess() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.EnrollRepo)
	mockUserId := uint64(123)
	mockCourseId := uint64(456)

	mockRepo.On("EnrollUser", mockUserId, mockCourseId).Return(nil)

	// Test
	underTest := services.NewEnrollService(mockRepo)
	err := underTest.EnrollUser(mockUserId, mockCourseId)

	// Assert
	is.NoError(err)
}

func (suite *EnrollServiceTestSuite) TestEnrollUserWhenRepoFails() {
	is := assert.New(suite.T())

	// Arrange
	mockRepo := new(mockRepositories.EnrollRepo)
	mockUserId := uint64(123)
	mockCourseId := uint64(456)

	mockRepo.On("EnrollUser", mockUserId, mockCourseId).Return(fmt.Errorf("repository error"))

	// Test
	underTest := services.NewEnrollService(mockRepo)
	err := underTest.EnrollUser(mockUserId, mockCourseId)

	// Assert
	is.NotNil(err)
	is.Equal("repository error", err.Error())
}

func TestEnrollService(t *testing.T) {
	suite.Run(t, new(EnrollServiceTestSuite))
}
