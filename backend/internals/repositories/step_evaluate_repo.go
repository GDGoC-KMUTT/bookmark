package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type stepEvaluateRepository struct {
	db *gorm.DB
}

func NewStepEvaluateRepository(db *gorm.DB) StepEvaluateRepository {
	return &stepEvaluateRepository{
		db: db,
	}
}

func (r *stepEvaluateRepository) GetStepEvalByStepId(stepId *uint64) ([]*models.StepEvaluate, error) {
	stepEvals := make([]*models.StepEvaluate, 0)

	result := r.db.Find(&stepEvals, "step_id = ?", stepId)
	if result.Error != nil {
		return nil, result.Error
	}
	return stepEvals, nil
}
