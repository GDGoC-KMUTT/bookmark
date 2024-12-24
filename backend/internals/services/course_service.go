package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type courseService struct {
	courseRepo    repositories.CourseRepository
	fieldTypeRepo repositories.FieldTypeRepository
}

func NewCourseService(courseRepo repositories.CourseRepository, fieldTypeRepo repositories.FieldTypeRepository) CourseService {
	return &courseService{
		courseRepo:    courseRepo,
		fieldTypeRepo: fieldTypeRepo,
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

func (r *courseService) GetAllFieldTypes() ([]payload.FieldType, error) {
	fieldTypes, tx := r.fieldTypeRepo.FindAllFieldType()
	if tx != nil {
		return nil, tx
	}

	var result []payload.FieldType
	for _, fieldType := range fieldTypes {
		result = append(result, payload.FieldType{
			Id:       fieldType.Id,
			Name:     fieldType.Name,
			ImageUrl: fieldType.ImageUrl,
		})
	}
	return result, nil

}
