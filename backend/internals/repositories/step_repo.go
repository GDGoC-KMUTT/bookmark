package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type StepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) StepRepository {
	return StepRepository{
		db: db,
	}
}

func (r *StepRepository) FindStepsByModuleID(moduleId string) ([]models.Step, error) {
	var steps []models.Step
	err := r.db.Where("module_id = ?", moduleId).Find(&steps).Error
	if err != nil {
		return nil, err
	}
	return steps, nil
}
