package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
)

type StepService interface {
	GetGems(stepId *uint64, userId *float64) (*int, *int, error)
	GetStepComment(stepId *uint64, userId *uint64) ([]payload.StepCommentInfo, error)
	CreteStpComment(stepId *uint64, userId *float64, content *string) error
	CreateOrDeleteStepCommentUpVote(userId *float64, stepCommentId *uint64) error
	GetStepInfo(stepId *uint64) (*payload.StepInfo, error)
	GetStepEvalInfo(stepId *uint64, userId *float64) ([]*payload.StepEvalInfo, error)
	CreateFileFormat(stepId *uint64, stepEvalId *uint64, userId *float64) (*string, error)
	CreateUserEval(payload *payload.CreateUserEvalReq) (*uint64, error)
	CheckStepEvalStatus(userEvalIds []*uint64, userId *uint64) ([]models.UserEvaluate, error)
	SubmitStepEvalTypeCheck(stepEvalId *uint64, userId *uint64) (*uint64, error)
}
