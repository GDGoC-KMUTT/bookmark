package repositories

import (
	"backend/internals/db/models"

	"gorm.io/gorm"
)

type stepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) StepRepo {
	return &stepRepository{
		db: db,
	}
}

// Use value receiver to match the StepRepo interface
func (r *stepRepository) FindStepsByModuleID(moduleId string) ([]models.Step, error) {
	var steps []models.Step
	err := r.db.Where("module_id = ?", moduleId).Find(&steps).Error
	if err != nil {
		return nil, err
	}
	return steps, nil
}
