package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type courseService struct {
	courseRepo repositories.CourseRepository
}

func NewCourseService(courseRepo repositories.CourseRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
	}
}

func (r *courseService) GetCourseByFieldId(fieldId *uint) ([]payload.CourseWithFieldType, error) {
	courses, fieldType, tx := r.courseRepo.FindCourseByFieldId(fieldId)
	if tx != nil {
		return nil, tx
	}
	var result []payload.CourseWithFieldType

	for _, course := range courses {
		result = append(result, payload.CourseWithFieldType{
			Id:         course.Id,
			Name:       course.Name,
			FieldId:    fieldType.Id,
			FieldName:  fieldType.Name,
			FieldImage: fieldType.ImageUrl,
		})
	}

	return result, nil
}
