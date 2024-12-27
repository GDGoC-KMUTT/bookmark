package services

import (
	"backend/internals/entities/response"
	"backend/internals/repositories"
)

type EnrollService interface {
	GetEnrollmentsByUserID(userId *string) (response.EnrollmentListResponse, error)
}

type enrollService struct {
	enrollRepo repositories.EnrollRepository
}

func NewEnrollService(enrollRepo repositories.EnrollRepository) EnrollService {
	return &enrollService{
		enrollRepo: enrollRepo,
	}
}

func (s *enrollService) GetEnrollmentsByUserID(userId *string) (response.EnrollmentListResponse, error) {
	enrollments, err := s.enrollRepo.FindEnrollmentsByUserID(userId)
	if err != nil {
		return response.EnrollmentListResponse{Enrollments: []response.EnrollmentListDTO{}}, err
	}

	var result []response.EnrollmentListDTO
	for _, e := range enrollments {
		// Fetch total steps and evaluated steps
		totalSteps, err := s.enrollRepo.GetTotalStepsByCourseID(*e.CourseId)
		if err != nil {
			totalSteps = 0 // Handle gracefully if error occurs
		}

		evaluatedSteps, err := s.enrollRepo.GetEvaluatedStepsByUserAndCourse(*e.UserId, *e.CourseId)
		if err != nil {
			evaluatedSteps = 0 // Handle gracefully if error occurs
		}

		// Calculate progress
		progress := 0.0
		if totalSteps > 0 {
			progress = (float64(evaluatedSteps) / float64(totalSteps)) * 100
		}

		// Add to result with progress
		result = append(result, response.EnrollmentListDTO{
			Id:         *e.Id,
			UserID:     *e.UserId,
			CourseID:   *e.CourseId,
			CourseName: *e.Course.Name,
			Progress:   progress,
			CreatedAt:  e.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:  e.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return response.EnrollmentListResponse{
		Enrollments: result,
	}, nil
}
