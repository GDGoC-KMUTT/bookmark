package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type userEvaluateRepo struct {
	db *gorm.DB
}

func NewUserEvaluateRepo(db *gorm.DB) UserEvaluateRepository {
	return &userEvaluateRepo{
		db: db,
	}
}

func (r *userEvaluateRepo) GetUserEvalByStepEvalId(stepEvalId *uint64, userId *float64) (*models.UserEvaluate, error) {
	userEval := new(models.UserEvaluate)

	result := r.db.Find(&userEval, "step_evaluate_id = ? AND user_id = ?", stepEvalId, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return userEval, nil
}
