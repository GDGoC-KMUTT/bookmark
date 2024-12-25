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

func (r *stepService) GetGems(stepId *uint64, userId *float64) (*int, *int, error) {
	stepEvals, err := r.stepEvalRepo.GetStepEvalByStepId(stepId)
	if err != nil {
		return nil, nil, err
	}

	totalGems := 0
	currentGems := 0
	for _, eval := range stepEvals {
		totalGems += *eval.Gem
		userEvals, err := r.userEvalRepo.GetUserEvalByStepEvalId(eval.Id, userId)
		if err != nil {
			return nil, nil, err
		}

		if *userEvals.Pass {
			currentGems += *eval.Gem
		}

	}

	return &totalGems, &currentGems, nil

}
