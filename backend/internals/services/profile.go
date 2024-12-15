package services

import "backend/internals/entities/payload"

type ProfileService interface {
	GetUserInfo(userId *string) (*payload.Profile, error)
}
