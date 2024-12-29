package services

import (
	"backend/internals/repositories"
)

type EnrollService struct {
    enrollRepo *repositories.EnrollRepository // pointer to EnrollRepository
}

func NewEnrollService(enrollRepo *repositories.EnrollRepository) *EnrollService {
    return &EnrollService{
        enrollRepo: enrollRepo,
    }
}

// EnrollUser enrolls a user in a course
func (s *EnrollService) EnrollUser(userId uint64, courseId uint64) error {
	// Use the EnrollUser method from the repository to enroll the user
	err := s.enrollRepo.EnrollUser(userId, courseId)
	if err != nil {
		return err
	}

	return nil
}
