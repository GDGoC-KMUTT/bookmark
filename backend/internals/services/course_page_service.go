package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"strconv"
)

type CoursePageRepo interface {
	FindCoursePageInfoByCoursePageID(coursePageId string) (*models.Course, error)
	FindCoursePageContentByCoursePageID(coursePageId string) ([]models.CourseContent, error)
}

type CourseRepo interface {
	FindCoursesByFieldId(fieldId uint64) ([]models.Course, error)
	FindFieldByFieldId(fieldId *uint64) (*models.FieldType, error)
}

type coursePageService struct {
	coursePageRepo CoursePageRepo
	courseRepo     CourseRepo
}

func NewCoursePageService(coursePageRepo CoursePageRepo, courseRepo CourseRepo) CoursePageServices {
	return &coursePageService{
		coursePageRepo: coursePageRepo,
		courseRepo:     courseRepo,
	}
}

func (s *coursePageService) GetCoursePageInfo(coursePageId string) (*payload.CoursePage, error) {
	coursePageEntity, err := s.coursePageRepo.FindCoursePageInfoByCoursePageID(coursePageId)
	if err != nil {
		return nil, err
	}

	return &payload.CoursePage{
		Id:      *coursePageEntity.Id,
		Name:    *coursePageEntity.Name,
		FieldId: *coursePageEntity.FieldId,
		Field:   coursePageEntity.Field.Name,
	}, nil
}

func (s *coursePageService) GetCoursePageContent(coursePageId string) ([]payload.CoursePageContent, error) {
	contentEntities, err := s.coursePageRepo.FindCoursePageContentByCoursePageID(coursePageId)
	if err != nil {
		return nil, err
	}

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

func (r *coursePageService) GetSuggestCourseByFieldId(fieldId string) ([]payload.SuggestCourse, error) {
	// fmt.Printf("Starting GetSuggestCourseByFieldId with fieldId: %s\n", fieldId)

	fieldIdUint64, err := strconv.ParseUint(fieldId, 10, 64)
	if err != nil {
		// fmt.Printf("Error parsing fieldId: %s, error: %v\n", fieldId, err)
		return nil, err
	}

	// fmt.Printf("Parsed fieldId to uint64: %d\n", fieldIdUint64)

	suggestCourses, err := r.courseRepo.FindCoursesByFieldId(fieldIdUint64)
	if err != nil {
		// fmt.Printf("Error fetching courses for fieldId %d: %v\n", fieldIdUint64, err)
		return nil, err
	}

	// fmt.Printf("Fetched %d courses for fieldId %d\n", len(suggestCourses), fieldIdUint64)

	var result []payload.SuggestCourse
	for _, course := range suggestCourses {
		// fmt.Printf("Processing course: %+v\n", course)

		field, err := r.courseRepo.FindFieldByFieldId(course.FieldId)
		if err != nil {
			// fmt.Printf("Error fetching field for course FieldId %d: %v\n", *course.FieldId, err)
			return nil, err
		}

		if field == nil {
			// fmt.Printf("Field is nil for FieldId: %d. Skipping course.\n", *course.FieldId)
			continue
		}

		// fmt.Printf("Field fetched for FieldId %d: %+v\n", *course.FieldId, field)

		result = append(result, payload.SuggestCourse{
			Id:            *course.Id,
			Name:          *course.Name,
			FieldId:       *course.FieldId,
			FieldName:     field.Name,
			FieldImageUrl: field.ImageUrl,
		})
	}

	// fmt.Printf("Returning %d suggested courses\n", len(result))
	return result, nil
}
