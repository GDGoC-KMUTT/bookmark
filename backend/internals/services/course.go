package services

import "backend/internals/entities/payload"

type CourseService interface {
	GetCourseByFieldId(field_id *uint) ([]payload.CourseWithFieldType, error)
	GetAllFieldTypes() ([]payload.FieldType, error)
	GetCurrentCourse(userID uint) (*payload.Course, error)
	GetTotalStepsByCourseId(courseID uint) (*payload.TotalStepsByCourseIdPayload, error)
}
