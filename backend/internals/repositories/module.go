package repositories

import "backend/internals/db/models"

type ModulesRepository interface {
	GetModuleById(moduleId *uint64) (*models.Module, error)
}
