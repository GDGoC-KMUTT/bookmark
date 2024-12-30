package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
	"fmt"
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

func (r *userEvaluateRepo) GetUserEvalByIdAndUserId(userEvalId *uint64, userId *uint64) (*models.UserEvaluate, error) {
	userEval := new(models.UserEvaluate)

	result := r.db.Find(&userEval, "id = ? AND user_id = ?", userEvalId, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
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

func (r *userEvaluateRepo) FindStepEvaluateIDsByStepID(stepID uint64) ([]uint64, error) {
	var stepEvaluateIDs []uint64

	// Query the step_evaluates table to get all IDs for the given step_id
	err := r.db.Model(&models.StepEvaluate{}).
		Where("step_id = ?", stepID).
		Pluck("id", &stepEvaluateIDs).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch step_evaluate IDs: %w", err)
	}

	return stepEvaluateIDs, nil
}

func (r *userEvaluateRepo) FindUserPassedEvaluateIDs(userID uint, stepID uint64) ([]uint64, error) {
	var userPassedIDs []uint64

	// Query the user_evaluates table to get IDs where pass is not null for the given user and step_id
	err := r.db.Table("user_evaluates").
		Select("step_evaluate_id").
		Where("user_id = ? AND step_evaluate_id IN (?) AND pass IS NOT NULL", userID,
			r.db.Table("step_evaluates").Select("id").Where("step_id = ?", stepID),
		).
		Scan(&userPassedIDs).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch user passed step_evaluate IDs: %w", err)
	}

	return userPassedIDs, nil
}

