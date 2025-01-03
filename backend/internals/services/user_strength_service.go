package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"backend/internals/utils"
	"fmt"
	"log"
	"strconv"
)

type userStrengthService struct {
	userStrengthRepo repositories.UserStrengthRepository
	fieldTypeRepo    repositories.FieldTypeRepository
	userRepo         repositories.UserRepository
}

func NewUserStrengthService(userStrengthRepo repositories.UserStrengthRepository, fieldTypeRepo repositories.FieldTypeRepository, userRepo repositories.UserRepository) UserStrengthService {
	return &userStrengthService{
		userStrengthRepo: userStrengthRepo,
		fieldTypeRepo:    fieldTypeRepo,
		userRepo:         userRepo,
	}
}

func (s *userStrengthService) GetStrengthDataByUserID(userId uint64) (*payload.StrengthDataResponse, error) {
	strengthFieldData, err := s.userStrengthRepo.GetStrengthDataByUserID(userId)
	if err != nil {
		log.Printf("Error in service while fetching strength data for user %d: %v", userId, err)
		return nil, err
	}
	fieldTypes, err := s.fieldTypeRepo.FindAllFieldTypes()
	if err != nil {
		return nil, err
	}

	if strengthFieldData == nil {
		log.Printf("No passing evaluations found for user %d", userId)
		return &payload.StrengthDataResponse{}, nil
	}

	strengthField := make([]string, 0)
	for _, strength := range strengthFieldData {
		strengthField = append(strengthField, strength.FieldName)
	}

	for _, item := range fieldTypes {
		if !utils.Contains(strengthField, *item.Name) {
			strengthFieldData = append(strengthFieldData, payload.StrengthFieldData{
				FieldName: *item.Name,
				TotalGems: int64(0),
			})
		}
	}

	user, err := s.userRepo.FindUserByID(utils.Ptr(strconv.FormatUint(userId, 10)))
	if err != nil {
		return nil, err
	}

	strengthData := &payload.StrengthDataResponse{
		Data:     strengthFieldData,
		Username: fmt.Sprintf("%s %s", *user.Firstname, *user.Lastname),
	}

	return strengthData, nil
}

func (s *userStrengthService) GetSuggestionCourse(userId uint64) ([]payload.CourseResponse, error) {
	courses, err := s.userStrengthRepo.GetSuggestionCourse(userId)
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
