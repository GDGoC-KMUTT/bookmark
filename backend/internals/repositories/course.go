package repositories

import "backend/internals/db/models"

type CourseRepository interface {
	FindCourseByFieldId(fieldId uint) ([]models.Course, *models.FieldType, error)
	GetCurrentCourse(userID uint) (*models.Course, error)
	GetTotalStepsByCourseId(courseID uint) (int, error)
	GetAllCourseSteps(courseID uint) ([]models.Step, error)
	FindEnrollCourseByUserId(userId int) ([]*models.Enroll, error)
	FindCourseByCourseId(courseId *uint64) (*models.Course, error)
	FindFieldByFieldId(fieldId *uint64) (*models.FieldType, error)
	FindCoursesByFieldId(fieldId uint64) ([]models.Course, error)
}
