package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type stepCommentUpVoteRepo struct {
	db *gorm.DB
}

func NewStepCommentUpVote(db *gorm.DB) StepCommentUpVoteRepository {
	return &stepCommentUpVoteRepo{
		db: db,
	}
}

func (r *stepCommentUpVoteRepo) GetStepCommentUpVoteByStepCommentId(stepCommentId *uint64) ([]*models.StepCommentUpvote, error) {
	stepCommentUpVote := make([]*models.StepCommentUpvote, 0)

	result := r.db.Find(&stepCommentUpVote, "step_comment_id = ?", stepCommentId)
	if result.Error != nil {
		return nil, result.Error
	}

	return stepCommentUpVote, nil
}
