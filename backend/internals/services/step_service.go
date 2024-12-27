package services

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"backend/internals/utils"
	"strconv"
)

type stepService struct {
	stepEvalRepo          repositories.StepEvaluateRepository
	userEvalRepo          repositories.UserEvaluateRepository
	userRepo              repositories.UserRepository
	stepCommentRepo       repositories.StepCommentRepository
	stepCommentUpVoteRepo repositories.StepCommentUpVoteRepository
}

func NewStepService(stepEvalRepo repositories.StepEvaluateRepository, userEvalRepo repositories.UserEvaluateRepository, userRepo repositories.UserRepository, stepCommentRepo repositories.StepCommentRepository, stepCommentUpVoteRepo repositories.StepCommentUpVoteRepository) StepService {
	return &stepService{
		stepEvalRepo:          stepEvalRepo,
		userEvalRepo:          userEvalRepo,
		userRepo:              userRepo,
		stepCommentRepo:       stepCommentRepo,
		stepCommentUpVoteRepo: stepCommentUpVoteRepo,
	}
}

func (r *stepService) GetGems(stepId *uint64, userId *float64) (*int, *int, error) {
	stepEvals, err := r.stepEvalRepo.GetStepEvalByStepId(stepId)
	if err != nil {
		return nil, nil, err
	}

	totalGems := 0
	currentGems := 0
	for _, eval := range stepEvals {
		totalGems += *eval.Gem
		userEvals, err2 := r.userEvalRepo.GetUserEvalByStepEvalId(eval.Id, userId)
		if err2 != nil {
			return nil, nil, err2
		}

		if userEvals == nil {
			continue
		}

		if *userEvals.Pass == true {
			currentGems += *eval.Gem
		}

	}

	return &totalGems, &currentGems, nil
}

func (r *stepService) GetStepComment(stepId *uint64) ([]payload.StepCommentInfo, error) {
	stepComments, err := r.stepCommentRepo.GetStepCommentByStepId(stepId)
	if err != nil {
		return nil, err
	}

	stepCommentInfo := make([]payload.StepCommentInfo, 0)
	for _, comment := range stepComments {
		user, err := r.userRepo.FindUserByID(utils.Ptr(strconv.FormatUint(*comment.UserId, 10)))
		if err != nil {
			return nil, err
		}

		stepCommentUpVote, err := r.stepCommentUpVoteRepo.GetStepCommentUpVoteByStepCommentId(comment.Id)
		if err != nil {
			return nil, err
		}

		stepCommentInfo = append(stepCommentInfo, payload.StepCommentInfo{
			StepCommentId: comment.Id,
			UserInfo: &payload.CommentedBy{
				UserId:    user.Id,
				FirstName: user.Firstname,
				Lastname:  user.Lastname,
				Email:     user.Email,
				PhotoUrl:  user.PhotoUrl,
			},
			Comment: comment.Content,
			UpVote:  utils.Ptr(len(stepCommentUpVote)),
		})
	}

	return stepCommentInfo, nil
}

func (r *stepService) CreteStpComment(stepId *uint64, userId *float64, content *string) error {
	stepComment := &models.StepComment{
		Content: content,
		StepId:  stepId,
		UserId:  utils.Ptr(uint64(*userId)),
	}

	if err := r.stepCommentRepo.CreateStepComment(stepComment); err != nil {
		return err
	}

	return nil
}

func (r *stepService) CreateStepCommentUpVote(userId *float64, stepCommentId *uint64) error {
	stepCommentUpVote := &models.StepCommentUpvote{
		StepCommentId: stepCommentId,
		UserId:        utils.Ptr(uint64(*userId)),
	}

	if err := r.stepCommentUpVoteRepo.CreateStepCommentUpVote(stepCommentUpVote); err != nil {
		return err
	}

	return nil
}
