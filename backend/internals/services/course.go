package services

import "backend/internals/entities/payload"

type CourseService interface {
	GetCourseByFieldId(field_id *uint) ([]payload.CourseWithFieldType, error)
	GetEnrollCourseByUserId(userId int) ([]*payload.EnrollwithCourse, error)
}