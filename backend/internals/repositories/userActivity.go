package repositories

import "backend/internals/db/models"

type UserActivityRepository interface {
	GetRecentActivitiesByUserID(userId *string) ([]models.UserActivity, error)
}
