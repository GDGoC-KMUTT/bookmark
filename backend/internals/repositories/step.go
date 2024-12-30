package repositories

import "backend/internals/db/models"

type StepRepository interface {
	GetStepById(stepId *uint64) (*models.Step, error)
	GetModuleIdByStepId(stepId *uint64) (*uint64, error)
	FindStepsByModuleID(moduleId *string) ([]*models.Step, error)
}
