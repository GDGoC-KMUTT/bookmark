package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type moduleService struct {
	moduleRepo repositories.ModuleRepo
}

func NewModuleService(moduleRepo repositories.ModuleRepo) ModuleServices {
	return &moduleService{
		moduleRepo: moduleRepo,
	}
}

func (s *moduleService) GetModuleInfo(moduleId string) (*payload.ModuleResponse, error) {
	// Fetch Module from the repository
	moduleEntity, err := s.moduleRepo.FindModuleInfoByModuleID(moduleId)
	if err != nil {
		return nil, err
	}

	// Handle nil entity
	if moduleEntity == nil {
		return nil, nil
	}

	// Map to payload.ModuleResponse
	return &payload.ModuleResponse{
		Id:          *moduleEntity.Id,
		Title:       *moduleEntity.Title,
		Description: moduleEntity.Description,
		ImageUrl:    moduleEntity.ImageUrl,
	}, nil
}
