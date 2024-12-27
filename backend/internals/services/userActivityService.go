package services

import (
	"backend/internals/entities/response"
	"backend/internals/repositories"
)

type UserActivityService interface {
	GetRecentActivityByUserID(userId *string) (*response.UserActivityResponse, error)
}

type userActivityService struct {
	userActivityRepo repositories.UserActivityRepository
}

func NewUserActivityService(userActivityRepo repositories.UserActivityRepository) UserActivityService {
	return &userActivityService{
		userActivityRepo: userActivityRepo,
	}
}

// GetRecentActivityByUserID fetches the most recent activity and returns it in a response-friendly format
func (s *userActivityService) GetRecentActivityByUserID(userId *string) (*response.UserActivityResponse, error) {
	// Fetch the most recent user activity
	activity, err := s.userActivityRepo.GetRecentActivityByUserID(userId)
	if err != nil {
		return nil, err
	}

	// Prepare the response structure
	return &response.UserActivityResponse{
		StepID:      *activity.StepId,
		StepTitle:   *activity.Step.Title,
		ModuleTitle: *activity.Step.Module.Title,
		CreatedAt:   activity.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   activity.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}
