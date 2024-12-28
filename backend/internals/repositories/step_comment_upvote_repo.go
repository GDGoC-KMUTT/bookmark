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

func (r *stepCommentUpVoteRepo) CreateStepCommentUpVote(stepCommentUpVote *models.StepCommentUpvote) error {
	return r.db.Create(stepCommentUpVote).Error
}

func (r *stepCommentUpVoteRepo) GetStepCommentUpVoteByStepCommentIdAndUserId(stepCommentId *uint64, userId *uint64) (*models.StepCommentUpvote, error) {
	stepCommentUpVote := new(models.StepCommentUpvote)

	result := r.db.Find(&stepCommentUpVote, "step_comment_id = ? AND user_id = ?", stepCommentId, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return stepCommentUpVote, nil
}

func (r *stepCommentUpVoteRepo) DeleteStepCommentUpVote(stepCommentId *uint64, userId *uint64) error {
	commentUpVote := new(models.StepCommentUpvote)
	result := r.db.Delete(&commentUpVote, "step_comment_id = ? AND user_id = ?", stepCommentId, userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
