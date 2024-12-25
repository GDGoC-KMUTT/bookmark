package repositories

import "backend/internals/db/models"

type UserEvaluateRepository interface {
	GetUserEvalByUserId(userId *uint64) ([]*models.UserEvaluate, error)
}
