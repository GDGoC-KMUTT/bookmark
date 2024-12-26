package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type stepCommentRepo struct {
	db *gorm.DB
}

func NewStepCommentRepository(db *gorm.DB) StepCommentRepository {
	return &stepCommentRepo{
		db: db,
	}
}

func (r *stepCommentRepo) GetStepCommentByStepId(stepId *uint64) ([]*models.StepComment, error) {
	stepComments := make([]*models.StepComment, 0)

	result := r.db.Find(&stepComments, "step_id = ?", stepId)
	if result.Error != nil {
		return nil, result.Error
	}

	return stepComments, nil
}
