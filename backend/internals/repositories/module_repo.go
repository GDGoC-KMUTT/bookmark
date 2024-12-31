package repositories

import (
	"backend/internals/db/models"

	"gorm.io/gorm"
)


type moduleRepo struct {
	db *gorm.DB
}

func NewModuleRepository(db *gorm.DB) ModulesRepository {
	return &moduleRepo{
		db: db,
	}
}

func (r *moduleRepo) GetModuleById(moduleId *uint64) (*models.Module, error) {
	module := new(models.Module)

	if result := r.db.First(&module, moduleId); result.Error != nil {
		return nil, result.Error
	}

	return module, nil
}

// Use value receiver to match the ModuleRepo interface
func (r *moduleRepo) FindModuleInfoByModuleID(moduleId string) (*models.Module, error) {
    var module models.Module
    err := r.db.Where("id = ?", moduleId).First(&module).Error
    if err != nil {
        return nil, err
    }
    return &module, nil
}
