package services

import "backend/internals/entities/payload"

type StepService interface {
	GetGems(stepId *uint64, userId *float64) (*int, *int, error)
	GetStepComment(stepId *uint64) ([]payload.StepCommentInfo, error)
	CreteStpComment(stepId *uint64, userId *float64, content *string) error
	CreateStepCommentUpVote(userId *float64, stepCommentId *uint64) error
}
