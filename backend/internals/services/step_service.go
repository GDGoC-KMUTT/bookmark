package services

import "backend/internals/repositories"

type stepService struct {
	stepEvalRepo repositories.StepEvaluateRepository
	userEvalRepo repositories.UserEvaluateRepository
}

func NewStepService(stepEvalRepo repositories.StepEvaluateRepository, userEvalRepo repositories.UserEvaluateRepository) StepService {
	return &stepService{
		stepEvalRepo: stepEvalRepo,
		userEvalRepo: userEvalRepo,
	}
}

func (r *stepService) GetGems(stepId *uint64, userId *uint64) (*uint64, *uint64, error) {
	stepEvals, err := r.stepEvalRepo.GetStepEvalByStepId(stepId)
	if err != nil {
		return nil, nil, err
	}

	userEvals, err := r.userEvalRepo.GetUserEvalByUserId(userId)

}
