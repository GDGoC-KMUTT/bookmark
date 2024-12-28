package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type CoursePageService struct {
	coursePageRepo repositories.CoursePageRepository
}

func NewCoursePageService(coursePageRepo repositories.CoursePageRepository) *CoursePageService {
	return &CoursePageService{
		coursePageRepo: coursePageRepo,
	}
}

func (s *CoursePageService) GetCoursePageInfo(coursePageId string) (*payload.CoursePage, error) {
	// Fetch course page from the repository
	coursePageEntity, err := s.coursePageRepo.FindCoursePageInfoByCoursePageID(coursePageId)
	if err != nil {
		return nil, err
	}

	// Map to payload.CoursePage
	return &payload.CoursePage{
		Id:      *coursePageEntity.Id,
		Name:    *coursePageEntity.Name,
		FieldId: *coursePageEntity.FieldId,
		Field:   coursePageEntity.Field.Name, // Assuming Field has a `Name` field
	}, nil
}

func (s *CoursePageService) GetCoursePageContent(coursePageId string) ([]payload.CoursePageContent, error) {
	// Fetch course page content from the repository
	contentEntities, err := s.coursePageRepo.FindCoursePageContentByCoursePageID(coursePageId)
	if err != nil {
		return nil, err
	}

	// Map to []payload.CoursePageContent
	contents := make([]payload.CoursePageContent, 0, len(contentEntities))
	for _, content := range contentEntities {
		contents = append(contents, payload.CoursePageContent{
			CoursePageId: *content.CourseId,
			Order:        *content.Order,
			Type:         *content.Type,
			Text:         content.Text,
			ModuleId:     content.ModuleId,
		})
	}

	return contents, nil
}
