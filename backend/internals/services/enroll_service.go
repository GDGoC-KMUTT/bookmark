package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"time"
)

type enrollService struct {
	enrollRepo repositories.EnrollRepository
}

func NewEnrollService(enrollRepo repositories.EnrollRepository) EnrollService {
	return &enrollService{
		enrollRepo: enrollRepo,
	}
}

// EnrollUser enrolls a user in a course
func (s *enrollService) EnrollUser(userId uint, courseId uint64) error {
	// Use the EnrollUser method from the repository to enroll the user
	err := s.enrollRepo.EnrollUser(userId, courseId)
	if err != nil {
		return err
	}

	return nil
}

func (s *enrollService) GetEnrollmentsByUserID(userId *string) (payload.EnrollmentListResponse, error) {
	enrollments, err := s.enrollRepo.FindEnrollmentsByUserID(userId)
	if err != nil {
		return payload.EnrollmentListResponse{Enrollments: []payload.EnrollmentListDTO{}}, err
	}

	var result []payload.EnrollmentListDTO
	for _, e := range enrollments {
		totalSteps, err := s.enrollRepo.GetTotalStepsByCourseID(*e.CourseId)
		if err != nil {
			totalSteps = 0 // Handle gracefully if error occurs
		}

		evaluatedSteps, err := s.enrollRepo.GetEvaluatedStepsByUserAndCourse(*e.UserId, *e.CourseId)
		if err != nil {
			evaluatedSteps = 0 // Handle gracefully if error occurs
		}

		progress := 0.0
		if totalSteps > 0 {
			progress = (float64(evaluatedSteps) / float64(totalSteps)) * 100
		}

		result = append(result, payload.EnrollmentListDTO{
			Id:         *e.Id,
			UserID:     *e.UserId,
			CourseID:   *e.CourseId,
			CourseName: *e.Course.Name,
			Progress:   progress,
			CreatedAt:  e.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  e.UpdatedAt.Format(time.RFC3339),
		})
	}

	return payload.EnrollmentListResponse{
		Enrollments: result,
	}, nil
}
