package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type stepAuthorRepo struct {
	db *gorm.DB
}

func NewStepAuthorRepository(db *gorm.DB) StepAuthorRepository {
	return &stepAuthorRepo{
		db: db,
	}
}

func (r *stepAuthorRepo) GetStepAuthorByStepId(stepId *uint64) ([]*models.StepAuthor, error) {
	stepAuthors := make([]*models.StepAuthor, 0)

	result := r.db.Find(&stepAuthors, "step_id = ?", stepId)
	if result.Error != nil {
		return nil, result.Error
	}

	return stepAuthors, nil
}
