package repositories

import "backend/internals/db/models"

type StepCommentUpVoteRepository interface {
	GetStepCommentUpVoteByStepCommentId(stepCommentId *uint64) ([]*models.StepCommentUpvote, error)
}
