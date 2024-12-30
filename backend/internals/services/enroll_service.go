package services

import (
	"backend/internals/repositories"
)

type EnrollService struct {
	enrollRepo repositories.EnrollRepo // Use the interface here
}

func NewEnrollService(enrollRepo repositories.EnrollRepo) *EnrollService { // Accept the interface
	return &EnrollService{
		enrollRepo: enrollRepo,
	}
}

// EnrollUser enrolls a user in a course
func (s *EnrollService) EnrollUser(userId uint, courseId uint64) error {
	// Use the EnrollUser method from the repository to enroll the user
	err := s.enrollRepo.EnrollUser(userId, courseId)
	if err != nil {
		return err
	}

	return nil
}
