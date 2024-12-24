package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type ModuleService struct {
    moduleRepo repositories.ModuleRepository
}

func NewModuleService(moduleRepo repositories.ModuleRepository) *ModuleService {
    return &ModuleService{
        moduleRepo: moduleRepo,
    }
}

func (s *ModuleService) GetModuleInfo(moduleId string) (*payload.ModuleResponse, error) {
	// Fetch Module from the repository
	moduleEntity, err := s.moduleRepo.FindModuleInfoByModuleID(moduleId)
	if err != nil {
		return nil, err
	}

	// Map to payload.ModuleResponse
	return &payload.ModuleResponse{
		Id:          *moduleEntity.Id,
		Title:       *moduleEntity.Title,
		Description: moduleEntity.Description,
		ImageUrl:    moduleEntity.ImageUrl,
	}, nil
}
