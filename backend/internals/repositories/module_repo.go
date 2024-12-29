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
