package services

import "backend/internals/entities/payload"

type UserActivityService interface {
	GetRecentActivitiesByUserID(userId *string) (*payload.UserActivitiesResponse, error)
}
