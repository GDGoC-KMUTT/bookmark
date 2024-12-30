package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type ModuleStepService struct {
	moduleStepRepo repositories.StepRepo // Use the interface here
}

func NewModuleStepService(moduleStepRepo repositories.StepRepo) *ModuleStepService { // Update the constructor
	return &ModuleStepService{
		moduleStepRepo: moduleStepRepo,
	}
}

func (s *ModuleStepService) GetModuleSteps(moduleId string) ([]payload.ModuleStepResponse, error) {
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
			Id:    derefUint64(step.Id),
			Title: derefString(step.Title),
			Check: derefString(step.Check),
		})
	}

	return steps, nil
}

// Helper functions to safely dereference pointers
func derefUint64(ptr *uint64) uint64 {
	if ptr == nil {
		return 0 // Default value if pointer is nil
	}
	return *ptr
}

func derefString(ptr *string) string {
	if ptr == nil {
		return "" // Default value if pointer is nil
	}
	return *ptr
}
