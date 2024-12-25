package repositories

import "backend/internals/db/models"

type UserEvaluateRepository interface {
	GetUserEvalByStepEvalId(stepEvalId *uint64, userId *float64) (*models.UserEvaluate, error)
}
