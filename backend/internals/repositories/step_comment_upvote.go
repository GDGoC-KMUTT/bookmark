package repositories

import "backend/internals/db/models"

type StepCommentUpVoteRepository interface {
	GetStepCommentUpVoteByStepCommentId(stepCommentId *uint64) ([]*models.StepCommentUpvote, error)
	CreateStepCommentUpVote(stepCommentUpVote *models.StepCommentUpvote) error
	GetStepCommentUpVoteByStepCommentIdAndUserId(stepCommentId *uint64, userId *uint64) (*models.StepCommentUpvote, error)
	DeleteStepCommentUpVote(stepCommentId *uint64, userId *uint64) error
}
