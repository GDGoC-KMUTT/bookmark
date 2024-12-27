package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type stepRepo struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) StepRepository {
	return &stepRepo{
		db: db,
	}
}

func (r *stepRepo) GetStepById(stepId *uint64) (*models.Step, error) {
	step := new(models.Step)
	result := r.db.First(&step, stepId)
	return step, result.Error
}
