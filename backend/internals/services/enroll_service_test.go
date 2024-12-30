package services_test

import (
	"backend/internals/services"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnrollUserWhenSuccess(t *testing.T) {
	is := assert.New(t)

	// Arrange
	mockRepo := new(mockRepositories.EnrollRepo)
	mockUserId := uint(123)
	mockCourseId := uint64(456)

	mockRepo.On("EnrollUser", mockUserId, mockCourseId).Return(nil)

	// Test
	underTest := services.NewEnrollService(mockRepo)
	err := underTest.EnrollUser(mockUserId, mockCourseId)

	// Assert
	is.NoError(err)
	mockRepo.AssertCalled(t, "EnrollUser", mockUserId, mockCourseId)
}

func TestEnrollUserWhenRepoFails(t *testing.T) {
	is := assert.New(t)

	// Arrange
	mockRepo := new(mockRepositories.EnrollRepo)
	mockUserId := uint(123)
	mockCourseId := uint64(456)

	mockRepo.On("EnrollUser", mockUserId, mockCourseId).Return(fmt.Errorf("repository error"))

	// Test
	underTest := services.NewEnrollService(mockRepo)
	err := underTest.EnrollUser(mockUserId, mockCourseId)

	// Assert
	is.NotNil(err)
	is.Equal("repository error", err.Error())
	mockRepo.AssertCalled(t, "EnrollUser", mockUserId, mockCourseId)
}
