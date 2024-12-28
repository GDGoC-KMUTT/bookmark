package repositories

import "backend/internals/db/models"

type UserEvaluateRepository interface {
	GetUserEvalByStepEvalIdUserId(stepEvalId *uint64, userId *float64) (*models.UserEvaluate, error)
	CreateUserEval(userEval *models.UserEvaluate) (*models.UserEvaluate, error)
	GetUserEvalById(userEvalId *uint64) (*models.UserEvaluate, error)
	GetPassAllUserEvalByStepEvalId(stepEvalId *uint64) ([]*models.UserEvaluate, error)
}
