package services_test

import (
	"backend/internals/db/models"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CoursePageServiceTestSuite struct {
	suite.Suite
}

func TestCoursePageService(t *testing.T) {
	suite.Run(t, new(CoursePageServiceTestSuite))
}

func (suite *CoursePageServiceTestSuite) TestGetCoursePageInfoSuccess() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	coursePageId := "1"
	mockCoursePage := models.Course{
		Id:      utils.Ptr(uint64(1)),
		Name:    utils.Ptr("Sample Course"),
		FieldId: utils.Ptr(uint64(2)),
		Field: &models.FieldType{
			Name: utils.Ptr("Science"),
		},
	}

	// Mock repository responses
	mockCoursePageRepo.On("FindCoursePageInfoByCoursePageID", coursePageId).Return(&mockCoursePage, nil)

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetCoursePageInfo(coursePageId)

	// Assertions
	is.Nil(err)
	is.NotNil(result)
	is.Equal(*mockCoursePage.Id, result.Id)
	is.Equal(*mockCoursePage.Name, result.Name)
	is.Equal(*mockCoursePage.FieldId, result.FieldId)
	is.Equal(*mockCoursePage.Field.Name, *result.Field) // Fix: Dereference pointer
}

func (suite *CoursePageServiceTestSuite) TestGetCoursePageContentSuccess() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	coursePageId := "1"
	mockContents := []models.CourseContent{
		{
			CourseId: utils.Ptr(uint64(1)),
			Order:    utils.Ptr(int64(1)),
			Type:     utils.Ptr("text"),
			Text:     utils.Ptr("Content 1"),
		},
		{
			CourseId: utils.Ptr(uint64(1)),
			Order:    utils.Ptr(int64(2)),
			Type:     utils.Ptr("module"),
			ModuleId: utils.Ptr(uint64(101)),
		},
	}

	// Mock repository responses
	mockCoursePageRepo.On("FindCoursePageContentByCoursePageID", coursePageId).Return(mockContents, nil)

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetCoursePageContent(coursePageId)

	// Assertions
	is.Nil(err)
	is.NotNil(result)
	is.Len(result, 2)
	is.Equal(*mockContents[0].CourseId, result[0].CoursePageId)
	is.Equal(*mockContents[1].ModuleId, *result[1].ModuleId) // Fix: Dereference pointer
}

func (suite *CoursePageServiceTestSuite) TestGetSuggestCourseByFieldIdSuccess() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	fieldId := "2"
	mockCourses := []models.Course{
		{
			Id:      utils.Ptr(uint64(1)),
			Name:    utils.Ptr("Course 1"),
			FieldId: utils.Ptr(uint64(2)),
		},
	}
	mockField := &models.FieldType{
		Id:       utils.Ptr(uint64(2)),
		Name:     utils.Ptr("Science"),
		ImageUrl: utils.Ptr("http://example.com/image.png"),
	}

	// Mock repository responses
	mockCourseRepo.On("FindCoursesByFieldId", uint64(2)).Return(mockCourses, nil)
	mockCourseRepo.On("FindFieldByFieldId", utils.Ptr(uint64(2))).Return(mockField, nil)

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetSuggestCourseByFieldId(fieldId)

	// Assertions
	is.Nil(err)
	is.NotNil(result)
	is.Len(result, 1)
	is.Equal(*mockCourses[0].Id, result[0].Id)
	is.Equal(*mockCourses[0].Name, result[0].Name)
	is.Equal(*mockField.Name, *result[0].FieldName)
	is.Equal(*mockField.ImageUrl, *result[0].FieldImageUrl)
}

func (suite *CoursePageServiceTestSuite) TestGetCoursePageInfoError() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	coursePageId := "1"
	mockCoursePageRepo.On("FindCoursePageInfoByCoursePageID", coursePageId).Return(nil, errors.New("repository error"))

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetCoursePageInfo(coursePageId)

	// Assertions
	is.Nil(result)
	is.NotNil(err)
	is.EqualError(err, "repository error")
}

func (suite *CoursePageServiceTestSuite) TestGetCoursePageContentError() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	coursePageId := "1"
	mockCoursePageRepo.On("FindCoursePageContentByCoursePageID", coursePageId).Return(nil, errors.New("repository error"))

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetCoursePageContent(coursePageId)

	// Assertions
	is.Nil(result)
	is.NotNil(err)
	is.EqualError(err, "repository error")
}

func (suite *CoursePageServiceTestSuite) TestGetSuggestCourseByFieldIdParseUintError() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Invalid field ID
	fieldId := "invalid"

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetSuggestCourseByFieldId(fieldId)

	// Assertions
	is.Nil(result)
	is.NotNil(err)
	is.EqualError(err, "strconv.ParseUint: parsing \"invalid\": invalid syntax")
}

func (suite *CoursePageServiceTestSuite) TestGetSuggestCourseByFieldIdFindCoursesError() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	fieldId := "2"
	mockCourseRepo.On("FindCoursesByFieldId", uint64(2)).Return(nil, errors.New("repository error"))

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetSuggestCourseByFieldId(fieldId)

	// Assertions
	is.Nil(result)
	is.NotNil(err)
	is.EqualError(err, "repository error")
}

func (suite *CoursePageServiceTestSuite) TestGetSuggestCourseByFieldIdFindFieldError() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	fieldId := "2"
	mockCourses := []models.Course{
		{Id: utils.Ptr(uint64(1)), Name: utils.Ptr("Course 1"), FieldId: utils.Ptr(uint64(2))},
	}
	mockCourseRepo.On("FindCoursesByFieldId", uint64(2)).Return(mockCourses, nil)
	mockCourseRepo.On("FindFieldByFieldId", utils.Ptr(uint64(2))).Return(nil, errors.New("repository error"))

	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetSuggestCourseByFieldId(fieldId)

	// Assertions
	is.Nil(result)
	is.NotNil(err)
	is.EqualError(err, "repository error")
}

func (suite *CoursePageServiceTestSuite) TestGetSuggestCourseByFieldIdFieldNil() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	fieldId := "1"
	mockCourses := []models.Course{
		{
			Id:      utils.Ptr(uint64(1)),
			Name:    utils.Ptr("Course 1"),
			FieldId: utils.Ptr(uint64(2)), // This will be used to fetch the field
		},
	}

	// Mock repository responses
	mockCourseRepo.On("FindCoursesByFieldId", uint64(1)).Return(mockCourses, nil)
	mockCourseRepo.On("FindFieldByFieldId", mockCourses[0].FieldId).Return(nil, nil) // Mocking field as nil

	// Service setup
	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetSuggestCourseByFieldId(fieldId)

	// Assertions
	is.Nil(err)
	is.Len(result, 0)
}

func (suite *CoursePageServiceTestSuite) TestGetSuggestCourseByFieldIdFieldNotNil() {
	is := assert.New(suite.T())

	mockCoursePageRepo := new(mockRepositories.CoursePageRepo)
	mockCourseRepo := new(mockRepositories.CourseRepository)

	// Mock data
	fieldId := "1"
	mockCourses := []models.Course{
		{
			Id:      utils.Ptr(uint64(1)),
			Name:    utils.Ptr("Course 1"),
			FieldId: utils.Ptr(uint64(2)), // This will be used to fetch the field
		},
	}

	mockField := &models.FieldType{
		Id:       utils.Ptr(uint64(2)),
		Name:     utils.Ptr("Science"),
		ImageUrl: utils.Ptr("http://example.com/image.png"),
	}

	// Mock repository responses
	mockCourseRepo.On("FindCoursesByFieldId", uint64(1)).Return(mockCourses, nil)
	mockCourseRepo.On("FindFieldByFieldId", mockCourses[0].FieldId).Return(mockField, nil) // Return valid field

	// Service setup
	service := services.NewCoursePageService(mockCoursePageRepo, mockCourseRepo)

	// Execute
	result, err := service.GetSuggestCourseByFieldId(fieldId)

	// Assertions
	is.Nil(err)
	is.NotNil(result)
	is.Len(result, 1) // Expect one result since Field is valid
	is.Equal(result[0].FieldName, mockField.Name)
}
