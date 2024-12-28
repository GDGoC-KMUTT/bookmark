package services

import (
	"backend/internals/entities/response"
	"backend/internals/repositories"
	"log"
)

type UserStrengthService interface {
	GetStrengthDataByUserID(userId uint64) ([]response.StrengthDataResponse, error)
}

type userStrengthService struct {
	repo repositories.UserStrengthRepository
}

func NewUserStrengthService(repo repositories.UserStrengthRepository) UserStrengthService {
	return &userStrengthService{
		repo: repo,
	}
}

// GetStrengthDataByUserID ทำการดึงข้อมูลคะแนน strength ของ user จาก repository
func (s *userStrengthService) GetStrengthDataByUserID(userId uint64) ([]response.StrengthDataResponse, error) {
	// เรียกใช้ repository เพื่อดึงข้อมูล strength ตาม userId
	strengthData, err := s.repo.GetStrengthDataByUserID(userId)
	if err != nil {
		log.Printf("Error in service while fetching strength data for user %d: %v", userId, err)
		return nil, err
	}

	return strengthData, nil
}
