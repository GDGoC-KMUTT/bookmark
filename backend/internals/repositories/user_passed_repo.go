package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type userPassedRepo struct {
	db *gorm.DB
}

func NewUserPassedRepository(db *gorm.DB) UserPassedRepository {
	return &userPassedRepo{
		db: db,
	}
}

func (r *userPassedRepo) GetUserPassedByStepIdCourseIdModuleId(stepId *uint64, courseId *uint64, moduleId *uint64, userPassedType *string) ([]*models.UserPass, error) {
	userPassed := make([]*models.UserPass, 0)

	result := r.db.Find(&userPassed, "step_id = ? AND course_id = ? AND module_id = ? AND type = ?", stepId, courseId, moduleId, userPassedType)
	if result.Error != nil {
		return nil, result.Error
	}

	return userPassed, nil
}

func (r *userPassedRepo) CreateUserPassed(userPassed *models.UserPass) error {
	return r.db.Create(userPassed).Error
}
