package repositories

import "backend/internals/db/models"

type UserPassedRepository interface {
	GetUserPassedByStepIdCourseIdModuleId(stepId *uint64, courseId *uint64, moduleId *uint64, userPassedType *string) ([]*models.UserPass, error)
	CreateUserPassed(userPassed *models.UserPass) error
}
