package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type profileService struct {
	userRepo repositories.UserRepository
}

func NewProfileService(userRepo repositories.UserRepository) ProfileService {
	return &profileService{
		userRepo: userRepo,
	}
}

func (r *profileService) GetUserInfo(userId *string) (*payload.Profile, error) {
	user, tx := r.userRepo.FindUserByID(userId)
	if tx != nil {
		return nil, tx
	}

	result := &payload.Profile{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		PhotoUrl:  user.PhotoUrl,
	}

	return result, nil
}
