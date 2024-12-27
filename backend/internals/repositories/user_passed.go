package repositories

import "backend/internals/db/models"

type UserPassedRepository interface {
	GetUserPassedByStepIdCourseIdModuleId(stepId *uint64, courseId *uint64, moduleId *uint64) ([]*models.UserPass, error)
}
