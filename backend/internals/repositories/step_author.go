package repositories

import "backend/internals/db/models"

type StepAuthorRepository interface {
	GetStepAuthorByStepId(stepId *uint64) ([]*models.StepAuthor, error)
}
