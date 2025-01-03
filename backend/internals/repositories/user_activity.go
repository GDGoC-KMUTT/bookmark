package repositories

import "backend/internals/db/models"

type UserActivityRepository interface {
	UpdateUserActivity(userId uint64, stepId uint64) error
	GetRecentActivitiesByUserID(userId *string) ([]models.UserActivity, error)
}
