package repositories

import (
	"backend/internals/db/models"
	"backend/internals/utils"
	"log"
	"errors"
	"gorm.io/gorm"
)

type userActivityRepository struct {
	db *gorm.DB
}

func NewUserActivityRepository(db *gorm.DB) UserActivityRepository {
	return &userActivityRepository{
		db: db,
	}
}

func (repo *userActivityRepository) UpdateUserActivity(userId uint64, stepId uint64) error {
    var existingActivity models.UserActivity

    // Check if the activity already exists
    result := repo.db.Where("user_id = ? AND step_id = ?", userId, stepId).First(&existingActivity)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            // If no record exists, create a new one
            newActivity := models.UserActivity{
                UserId:    &userId,
                StepId:    &stepId,
                CreatedAt: utils.TimeNowPtr(),
                UpdatedAt: utils.TimeNowPtr(),
            }
            if err := repo.db.Create(&newActivity).Error; err != nil {
                return err
            }
        } else {
            // Return other errors encountered during the query
            return result.Error
        }
    } else {
        // If record exists, explicitly update the UpdatedAt field
        err := repo.db.Model(&existingActivity).
            Where("user_id = ? AND step_id = ?", userId, stepId).
            Update("updated_at", utils.TimeNowPtr()).Error
        if err != nil {
            return err
        }
    }

    return nil
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
