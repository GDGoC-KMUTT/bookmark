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

func (r *userEvaluateRepo) GetUserEvalByStepEvalIdUserId(stepEvalId *uint64, userId *float64) (*models.UserEvaluate, error) {
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

func (r *userEvaluateRepo) CreateUserEval(userEval *models.UserEvaluate) (*models.UserEvaluate, error) {
	result := r.db.Create(userEval)
	if result.Error != nil {
		return nil, result.Error
	}

	return userEval, nil
}

func (r *userEvaluateRepo) GetUserEvalById(userEvalId *uint64) (*models.UserEvaluate, error) {
	userEval := new(models.UserEvaluate)

	if err := r.db.First(&userEval, userEvalId); err != nil {
		return nil, err.Error
	}

	return userEval, nil
}

func (r *userEvaluateRepo) GetPassAllUserEvalByStepEvalId(stepEvalId *uint64) ([]*models.UserEvaluate, error) {
	userEval := make([]*models.UserEvaluate, 0)

	result := r.db.Find(&userEval, "step_evaluate_id = ? AND pass = true", stepEvalId)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return userEval, nil
}
