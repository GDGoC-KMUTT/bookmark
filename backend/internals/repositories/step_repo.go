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

// Use value receiver to match the StepRepo interface
func (r *stepRepo) FindStepsByModuleID(moduleId string) ([]models.Step, error) {
    var steps []models.Step
    err := r.db.Where("module_id = ?", moduleId).Find(&steps).Error
    if err != nil {
        return nil, err
    }
    return steps, nil
}

