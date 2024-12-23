package repositories

import "backend/internals/db/models"

type CourseRepository interface {
	FindCourseByFieldId(field_id *uint) ([]models.Course, *models.FieldType, error)
	FindEnrollCourseByUserId(userId int) ([]*models.Enroll, error)
}