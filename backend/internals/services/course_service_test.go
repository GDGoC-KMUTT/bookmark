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
	"testing"
)

type CourseServiceTestSuite struct {
	suite.Suite
}


func (suite *CourseServiceTestSuite) TestGetEnrollCourseByUserIdWhenSuccess() {
	is := assert.New(suite.T())
	mockCourseRepo := new(mockRepositories.CourseRepository)

	mockUserId := 1
	mockEnrollments := []*models.Enroll{
		{
			Id:       utils.Ptr(uint64(1)),
			UserId:   utils.Ptr(uint64(1)),
			CourseId: utils.Ptr(uint64(1)),
		},
	}

	mockCourse := &models.Course{
		Id:      utils.Ptr(uint64(1)),
		Name:    utils.Ptr("Course 1"),
		FieldId: utils.Ptr(uint64(1)),
	}

	mockField := &models.FieldType{
		Id:       utils.Ptr(uint64(1)),
		Name:     utils.Ptr("Field 1"),
		ImageUrl: utils.Ptr("image_url"),
	}

	mockCourseRepo.EXPECT().FindEnrollCourseByUserId(mock.Anything).Return(mockEnrollments, nil)
	mockCourseRepo.EXPECT().FindCourseByCourseId(mock.Anything).Return(mockCourse, nil)
	mockCourseRepo.EXPECT().FindFieldByFieldId(mock.Anything).Return(mockField, nil)

	underTest := services.NewCourseService(mockCourseRepo)

	result, err := underTest.GetEnrollCourseByUserId(mockUserId)

	is.NoError(err)
	is.NotNil(result)
	is.Equal(1, len(result))
	is.Equal("Course 1", *result[0].CourseName.Name)
	is.Equal("Field 1", *result[0].FieldName)
	is.Equal("image_url", *result[0].FieldImageURL)
}

func (suite *CourseServiceTestSuite) TestGetEnrollCourseByUserIdWhenError() {
	is := assert.New(suite.T())

	mockCourseRepo := new(mockRepositories.CourseRepository)
	mockUserId := 1

	mockCourseRepo.EXPECT().FindEnrollCourseByUserId(mock.Anything).Return(nil, fmt.Errorf("error fetching enrollments"))

	underTest := services.NewCourseService(mockCourseRepo)

	result, err := underTest.GetEnrollCourseByUserId(mockUserId)

	is.Error(err)
	is.Nil(result)
	is.Equal("error fetching enrollments", err.Error())
}

func TestCourseService(t *testing.T) {
	suite.Run(t, new(CourseServiceTestSuite))
}