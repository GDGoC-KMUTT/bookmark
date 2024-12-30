package repositories

import (
	"backend/internals/db/models"

	"gorm.io/gorm"
)

type ModuleRepository struct {
    db *gorm.DB
}

func NewModuleRepository(db *gorm.DB) ModuleRepository {
    return ModuleRepository{
        db: db,
    }
}

// Use value receiver to match the ModuleRepo interface
func (r ModuleRepository) FindModuleInfoByModuleID(moduleId string) (*models.Module, error) {
    var module models.Module
    err := r.db.Where("id = ?", moduleId).First(&module).Error
    if err != nil {
        return nil, err
    }
    return &module, nil
}
