package services

import (
	"backend/internals/repositories"
)

type userActivityService struct {
	userActivityRepo repositories.UserActivityRepository
}

func NewUserActivityService(userActivityRepo repositories.UserActivityRepository) UserActivityService {
	return &userActivityService{
		userActivityRepo: userActivityRepo,
	}
}

func (s *userActivityService) UpdateUserActivity(userId uint64, stepId uint64) error {
	err := s.userActivityRepo.UpdateUserActivity(userId, stepId)
	if err != nil {
		return err
	}

	return nil
}
