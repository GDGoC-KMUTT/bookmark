package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type gemService struct {
	gemRepo repositories.UserRepository
}

func NewGemService(gemRepo repositories.UserRepository) GemService {
	return &gemService{
		gemRepo: gemRepo,
	}
}

func (s *gemService) GetTotalGems(userID uint) (*payload.GemTotal, error) {
	totalGems, err := s.gemRepo.GetTotalGemsByUserID(userID)
	if err != nil {
		return nil, err
	}

	result := &payload.GemTotal{
		UserID: userID,
		Total:  totalGems,
	}
	return result, nil
}
