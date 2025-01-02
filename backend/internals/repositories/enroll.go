package repositories

import "backend/internals/db/models"

type EnrollRepository interface {
	EnrollUser(userId uint, courseId uint64) error
	FindEnrollmentsByUserID(userId *string) ([]models.Enroll, error)
	GetTotalStepsByCourseID(courseId uint64) (int64, error)
	GetEvaluatedStepsByUserAndCourse(userId uint64, courseId uint64) (int64, error)
}
