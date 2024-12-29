package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"log"
)

type userStrengthService struct {
	repo repositories.UserStrengthRepository
}

func NewUserStrengthService(repo repositories.UserStrengthRepository) UserStrengthService {
	return &userStrengthService{
		repo: repo,
	}
}

func (s *userStrengthService) GetStrengthDataByUserID(userId uint64) ([]payload.StrengthDataResponse, error) {
	strengthData, err := s.repo.GetStrengthDataByUserID(userId)
	if err != nil {
		log.Printf("Error in service while fetching strength data for user %d: %v", userId, err)
		return nil, err
	}

	if strengthData == nil {
		log.Printf("No passing evaluations found for user %d", userId)
		return []payload.StrengthDataResponse{}, nil
	}

	return strengthData, nil
}
func (s *userStrengthService) GetSuggestionCourse(userId uint64) ([]payload.CourseResponse, error) {
	courses, err := s.repo.GetSuggestionCourse(userId)
	if err != nil {
		log.Printf("Error in service while fetching random courses for user %d: %v", userId, err)
		return nil, err
	}

	// Transform to response DTOs
	var responses []payload.CourseResponse
	for _, course := range courses {
		if course.Id != nil && course.Name != nil && course.Field != nil && course.Field.Id != nil && course.Field.Name != nil {
			responses = append(responses, payload.CourseResponse{
				ID:   *course.Id,
				Name: *course.Name,
				Field: payload.FieldResponse{
					ID:       *course.Field.Id,
					Name:     *course.Field.Name,
					ImageUrl: course.Field.ImageUrl,
				},
			})
		}
	}

	return responses, nil
}
