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

func (r *stepRepo) GetModuleIdByStepId(stepId *uint64) (*uint64, error) {
	step := new(models.Step)
	result := r.db.First(&step, stepId)
	return step.ModuleId, result.Error
}

func (r *stepRepo) FindStepsByModuleID(moduleId *string) ([]*models.Step, error) {
	if moduleId == nil || *moduleId == "" {
		return nil, gorm.ErrInvalidData
	}

	var steps []*models.Step
	result := r.db.Where("module_id = ?", moduleId).Order("id ASC").Find(&steps)
	if result.Error != nil {
		return nil, result.Error
	}

	return steps, nil
}
