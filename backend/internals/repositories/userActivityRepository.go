package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
	"log"
)

type UserActivityRepository interface {
	GetRecentActivityByUserID(userId *string) (*models.UserActivity, error)
}

type userActivityRepository struct {
	db *gorm.DB
}

func NewUserActivityRepository(db *gorm.DB) UserActivityRepository {
	return &userActivityRepository{
		db: db,
	}
}

// GetRecentActivityByUserID fetches the most recent activity of a user
func (r *userActivityRepository) GetRecentActivityByUserID(userId *string) (*models.UserActivity, error) {
	var activity models.UserActivity
	// Query the most recent activity for the user, ordered by CreatedAt DESC
	err := r.db.
		Preload("Step"). // Preload related Step data
		Where("user_id = ?", userId).
		Order("created_at DESC"). // Ensure the most recent activity is selected
		Limit(1). // Limit the result to just the most recent activity
		First(&activity).Error

	if err != nil {
		log.Printf("Error fetching recent activity for user %s: %v", *userId, err)
		return nil, err
	}
	return &activity, nil
}
