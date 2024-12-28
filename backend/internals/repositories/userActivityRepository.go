package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
	"log"
)

type UserActivityRepository interface {
	GetRecentActivitiesByUserID(userId *string) ([]models.UserActivity, error)
}

type userActivityRepository struct {
	db *gorm.DB
}

func NewUserActivityRepository(db *gorm.DB) UserActivityRepository {
	return &userActivityRepository{
		db: db,
	}
}

func (r *userActivityRepository) GetRecentActivitiesByUserID(userId *string) ([]models.UserActivity, error) {
	var activities []models.UserActivity
	err := r.db.
		Preload("Step.Module").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Limit(10).
		Find(&activities).Error

	if err != nil {
		log.Printf("Error fetching recent activities for user %s: %v", *userId, err)
		return nil, err
	}

	return activities, nil
}
