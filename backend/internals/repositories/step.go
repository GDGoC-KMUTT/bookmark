package repositories

import "backend/internals/db/models"

type StepRepository interface {
	GetStepById(stepId *uint64) (*models.Step, error)
}
