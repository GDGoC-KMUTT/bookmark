package services

import (
	"backend/internals/entities/payload"
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

func (s *userActivityService) GetRecentActivitiesByUserID(userId *string) (*payload.UserActivitiesResponse, error) {
	activities, err := s.userActivityRepo.GetRecentActivitiesByUserID(userId)
	if err != nil {
		return nil, err
	}

	var activityResponses []payload.UserActivityResponse
	for _, activity := range activities {
		stepTitle := "Unknown Step"
		moduleTitle := "Unknown Module"

		if activity.Step != nil {
			stepTitle = *activity.Step.Title
			if activity.Step.Module != nil {
				moduleTitle = *activity.Step.Module.Title
			}
		}

		activityResponses = append(activityResponses, payload.UserActivityResponse{
			StepID:      *activity.StepId,
			StepTitle:   stepTitle,
			ModuleTitle: moduleTitle,
			CreatedAt:   activity.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   activity.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return &payload.UserActivitiesResponse{
		Activities: activityResponses,
	}, nil
}
