package repositories

import "backend/internals/db/models"


type ModulesRepository interface {
	FindModuleInfoByModuleID(moduleId string) (*models.Module, error)
	GetModuleById(moduleId *uint64) (*models.Module, error)
}
