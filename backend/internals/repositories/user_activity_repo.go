package repositories

import (
	"backend/internals/db/models"
	"errors"
	"gorm.io/gorm"
	"time"
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
                CreatedAt: timePtr(time.Now()),
                UpdatedAt: timePtr(time.Now()),
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
        err := repo.db.Model(&existingActivity).Where("user_id = ? AND step_id = ?", userId, stepId).
            Update("updated_at", timePtr(time.Now())).Error
        if err != nil {
            return err
        }
    }

    return nil
}
