package services

import (
	// "backend/internals/entities/response"
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type CourseService struct {
    courseRepo repositories.CourseRepository
}

func NewCourseService(courseRepo repositories.CourseRepository) *CourseService {
    return &CourseService{
        courseRepo: courseRepo,
    }
}

func (s *CourseService) GetCourseInfo(courseId string) (*payload.Course, error) {
	// Fetch course from the repository
	courseEntity, err := s.courseRepo.FindCourseInfoByCourseID(courseId)
	if err != nil {
		return nil, err
	}

	// Map to payload.Course
	return &payload.Course{
		Id:      *courseEntity.Id,
		Name:    *courseEntity.Name,
		FieldId: *courseEntity.FieldId,
		Field:   courseEntity.Field.Name, // Assuming Field has a `Name` field
	}, nil
}

func (s *CourseService) GetCourseContent(courseId string) ([]payload.CourseContent, error) {
	// Fetch course content from the repository
	contentEntities, err := s.courseRepo.FindCourseContentByCourseID(courseId)
	if err != nil {
		return nil, err
	}

	// Map to []payload.CourseContent
	contents := make([]payload.CourseContent, 0, len(contentEntities))
	for _, content := range contentEntities {
		contents = append(contents, payload.CourseContent{
			CourseId: *content.CourseId,
			Order:    *content.Order,
			Type:     *content.Type,
			Text:     content.Text,
			ModuleId: content.ModuleId,
		})
	}

	return contents, nil
}
