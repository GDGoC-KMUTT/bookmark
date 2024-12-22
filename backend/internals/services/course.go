package services

import "backend/internals/entities/payload"

type CourseService interface {
	GetCourseByFieldId(field_id *uint) ([]payload.CourseWithFieldType, error)
}
