package repositories

import "backend/internals/db/models"

type StepCommentRepository interface {
	GetStepCommentByStepId(stepId *uint64) ([]*models.StepComment, error)
}
