package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"fmt"
)

type moduleStepService struct {
	stepRepo       repositories.StepRepository
	userEvaluateRepo repositories.UserEvaluateRepository
}

func NewModuleStepService(stepRepo repositories.StepRepository, userEvaluateRepo repositories.UserEvaluateRepository) ModuleStepServices {
	return &moduleStepService{
		stepRepo:       stepRepo,
		userEvaluateRepo: userEvaluateRepo,
	}
}

func (s *moduleStepService) GetModuleSteps(userID uint, moduleID string) ([]payload.ModuleStepResponse, error) {
	// Fetch steps for the module
	steps, err := s.stepRepo.FindStepsByModuleID(&moduleID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch steps for module ID %s: %w", moduleID, err)
	}

	if len(steps) == 0 {
		return nil, fmt.Errorf("no steps found for module ID %s", moduleID)
	}

	// Prepare response
	var stepResponses []payload.ModuleStepResponse
	for _, step := range steps {
		// Validate step data
		if step.Id == nil || step.Title == nil {
			return nil, fmt.Errorf("invalid step data: missing ID or Title")
		}

		// Get all step_evaluate IDs for the current step
		stepEvaluateIDs, err := s.userEvaluateRepo.FindStepEvaluateIDsByStepID(*step.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch step evaluate IDs for step ID %d: %w", *step.Id, err)
		}

		// Get all step_evaluate IDs where user has passed
		userPassedIDs, err := s.userEvaluateRepo.FindUserPassedEvaluateIDs(userID, *step.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch user evaluations for step ID %d: %w", *step.Id, err)
		}

		// Determine the 'Check' status
		check := len(stepEvaluateIDs) > 0 && len(stepEvaluateIDs) == len(userPassedIDs)

		// Append the step response
		stepResponses = append(stepResponses, payload.ModuleStepResponse{
			Id:    *step.Id,
			Title: *step.Title,
			Check: check,
		})
	}

	return stepResponses, nil
}
