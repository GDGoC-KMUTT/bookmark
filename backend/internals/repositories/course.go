package repositories

import "backend/internals/db/models"

type CourseRepository interface {
	FindCourseByFieldId(field_id *uint) ([]models.Course, *models.FieldType, error)
	GetCurrentCourse(userID uint) (*models.Course, error)
	GetTotalStepsByCourseId(courseID uint) (int, error)
	GetAllCourseSteps(courseID uint) ([]models.Step, error)
}
