package repositories

import "backend/internals/db/models"

type StepRepo interface {
	FindStepsByModuleID(moduleId string) ([]models.Step, error)
}
