package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type userActivityService struct {
	userActivityRepo  repositories.UserActivityRepository
	stepRepo          repositories.StepRepository
	courseContentRepo repositories.CourseContentRepository
}

func NewUserActivityService(userActivityRepo repositories.UserActivityRepository, stepRepo repositories.StepRepository, courseContentRepo repositories.CourseContentRepository) UserActivityService {
	return &userActivityService{
		userActivityRepo:  userActivityRepo,
		stepRepo:          stepRepo,
		courseContentRepo: courseContentRepo,
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

		moduleId, err := s.stepRepo.GetModuleIdByStepId(activity.StepId)
		if err != nil {
			return nil, err
		}

		courseId, err := s.courseContentRepo.GetCourseIdByModuleId(moduleId)
		if err != nil {
			return nil, err
		}

		activityResponses = append(activityResponses, payload.UserActivityResponse{
			CourseId:    *courseId,
			ModuleId:    *moduleId,
			StepID:      *activity.StepId,
			StepTitle:   stepTitle,
			ModuleTitle: moduleTitle,
			CreatedAt:   activity.CreatedAt,
			UpdatedAt:   activity.UpdatedAt,
		})
	}

	return &payload.UserActivitiesResponse{
		Activities: activityResponses,
	}, nil
}

func (s *userActivityService) UpdateUserActivity(userId uint64, stepId uint64) error {
	err := s.userActivityRepo.UpdateUserActivity(userId, stepId)
	if err != nil {
		return err
	}

	return nil
}
