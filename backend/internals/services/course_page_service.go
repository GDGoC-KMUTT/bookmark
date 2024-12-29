package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"fmt"
	"strconv"
)

type CoursePageService struct {
	coursePageRepo repositories.CoursePageRepository
	courseRepo     repositories.CourseRepository
}

func NewCoursePageService(coursePageRepo repositories.CoursePageRepository, courseRepo repositories.CourseRepository) *CoursePageService {
	return &CoursePageService{
		coursePageRepo: coursePageRepo,
		courseRepo:     courseRepo,  // Initialize courseRepo
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

func (r *CoursePageService) GetSuggestCourseByFieldId(fieldId string) ([]payload.SuggestCourse, error) {
    // Convert fieldId to uint64 since your model uses *uint64 for FieldId
    fieldIdUint64, err := strconv.ParseUint(fieldId, 10, 64)
    if err != nil {
        fmt.Printf("Error converting fieldId: %s, Error: %v\n", fieldId, err)
        return nil, err
    }

    // Use courseRepo to get the suggest courses by fieldId
    suggestCourses, err := r.courseRepo.FindCoursesByFieldId(fieldIdUint64)
    if err != nil {
        fmt.Printf("Error fetching suggested courses for FieldId: %d, Error: %v\n", fieldIdUint64, err)
        return nil, err
    }
    if suggestCourses == nil {
        return []payload.SuggestCourse{}, nil
    }

    var result []payload.SuggestCourse
    for _, course := range suggestCourses {
        // Use courseRepo to get the field details by FieldId
        field, err := r.courseRepo.FindFieldByFieldId(course.FieldId)
        if err != nil {
            fmt.Printf("Error fetching field details for FieldId: %d, Error: %v\n", course.FieldId, err)
            return nil, err
        }

        if field == nil {
            fmt.Printf("Field details for FieldId: %d not found\n", course.FieldId)
            continue // Skip this course if field details are missing
        }

        // Append the result as a SuggestCourse
        result = append(result, payload.SuggestCourse{
            Id:          *course.Id,
            Name:        *course.Name,
            FieldId:     *course.FieldId,
            FieldName:   field.Name,  // Pass the pointer as required by the model
            FieldImageUrl: field.ImageUrl,  // Pass the pointer as required by the model
        })
    }

    return result, nil
}
