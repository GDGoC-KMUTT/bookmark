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

func (r *courseService) GetCurrentCourse(userID uint) (*payload.Course, error) {
	course, err := r.courseRepo.GetCurrentCourse(userID)
	if err != nil {
		return nil, err
	}

	courseDetails := &payload.Course{
		Id:   course.Id,
		Name: course.Name,
	}

	return courseDetails, nil
}

func (r *courseService) GetTotalStepsByCourseId(courseID uint) (*payload.TotalStepsByCourseIdPayload, error) {
	totalSteps, err := r.courseRepo.GetTotalStepsByCourseId(courseID)
	if err != nil {
		return nil, err
	}

	return &payload.TotalStepsByCourseIdPayload{
		CourseId:   courseID,
		TotalSteps: totalSteps,
	}, nil
}