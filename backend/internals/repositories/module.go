package repositories

import "backend/internals/db/models"

type ModuleRepo interface {
	FindModuleInfoByModuleID(moduleId string) (*models.Module, error)
}
