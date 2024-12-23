package services

import (
	"backend/internals/repositories"
	"fmt"
)

type progressService struct {
	userRepo   repositories.UserRepository
	courseRepo repositories.CourseRepository
}

func NewProgressService(userRepo repositories.UserRepository, courseRepo repositories.CourseRepository) ProgressService {
	return &progressService{
		userRepo: userRepo,
		courseRepo: courseRepo,
	}
}

// Ensure that progressService implements the ProgressService interface
func (s *progressService) GetCompletionPercentage(userID uint, courseID uint) (float64, error) {
    // Fetch all course steps
    steps, err := s.courseRepo.GetAllCourseSteps(courseID)
    if err != nil {
        return 0, fmt.Errorf("failed to fetch course steps: %w", err)
    }

    // Fetch user's completed steps
    userSteps, err := s.userRepo.GetUserCompletedSteps(userID)
    if err != nil {
        return 0, fmt.Errorf("failed to fetch user completed steps: %w", err)
    }

    completedCount := 0
	for _, step := range steps {
		for _, userStep := range userSteps {
			stepID := *step.Id
			userStepID := *userStep.StepId
						
			if stepID == userStepID {
				completedCount++
				break
			}
		}
	}	
	
    totalSteps := len(steps)
    if totalSteps == 0 {
        return 0, fmt.Errorf("no steps found for course ID %d", courseID)
    }

    percentage := float64(completedCount) / float64(totalSteps) * 100

    return percentage, nil
}
