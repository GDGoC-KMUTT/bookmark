package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type moduleStepService struct {
	moduleStepRepo repositories.StepRepo // Use the interface here
}

func NewModuleStepService(moduleStepRepo repositories.StepRepo) ModuleStepServices { // Update the constructor
	return &moduleStepService{
		moduleStepRepo: moduleStepRepo,
	}
}

func (s *moduleStepService) GetModuleSteps(moduleId string) ([]payload.ModuleStepResponse, error) {
	// Fetch Steps from the repository
	stepEntities, err := s.moduleStepRepo.FindStepsByModuleID(moduleId)
	if err != nil {
		return nil, err
	}

	// Handle nil or empty result
	if stepEntities == nil {
		return []payload.ModuleStepResponse{}, nil
	}

	// Map to a slice of payload.ModuleStepResponse
	var steps []payload.ModuleStepResponse
	for _, step := range stepEntities {
		steps = append(steps, payload.ModuleStepResponse{
			Id:    *step.Id,
			Title: *step.Title,
			Check: *step.Check,
		})
	}

	return steps, nil
}
