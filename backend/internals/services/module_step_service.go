package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"fmt"
)

type moduleStepService struct {
	stepRepo         repositories.StepRepository
	userEvaluateRepo repositories.UserEvaluateRepository
	userRepo         repositories.UserRepository
}

func NewModuleStepService(stepRepo repositories.StepRepository, userEvaluateRepo repositories.UserEvaluateRepository, userRepo repositories.UserRepository) ModuleStepServices {
	return &moduleStepService{
		stepRepo:         stepRepo,
		userEvaluateRepo: userEvaluateRepo,
		userRepo:         userRepo,
	}
}

func (s *moduleStepService) GetModuleSteps(userID uint, moduleID string) ([]payload.ModuleStep, error) {
	// Fetch steps for the module
	steps, err := s.stepRepo.FindStepsByModuleID(&moduleID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch steps for module ID %s: %w", moduleID, err)
	}

	if len(steps) == 0 {
		return nil, fmt.Errorf("no steps found for module ID %s", moduleID)
	}

	// Prepare response
	var stepResponses []payload.ModuleStep
	for _, step := range steps {
		// Validate step data
		if step.Id == nil || step.Title == nil {
			return nil, fmt.Errorf("invalid step data: missing ID or Title")
		}

		userPass, err := s.userRepo.GetUserPassByUserID(userID)
		if err != nil {
			return nil, err
		}
		
		// Determine the 'Check' status
		check := userPass > 0

		// Append the step response
		stepResponses = append(stepResponses, payload.ModuleStep{
			Id:    *step.Id,
			Title: *step.Title,
			Check: check,
		})
	}

	return stepResponses, nil
}
