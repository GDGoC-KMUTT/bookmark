package repositories

import (
	"backend/internals/db/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type EnrollRepository struct {
	db *gorm.DB
}

func NewEnrollRepository(db *gorm.DB) *EnrollRepository {
	return &EnrollRepository{
		db: db,
	}
}

// EnrollUser enrolls a user in a course
func (repo *EnrollRepository) EnrollUser(userId, courseId uint64) error {
	// Check if the user is already enrolled
	var existingEnrollment models.Enroll
	result := repo.db.Where("user_id = ? AND course_id = ?", userId, courseId).First(&existingEnrollment)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// If no enrollment found, create a new enrollment record
	if result.Error == gorm.ErrRecordNotFound {
		enrollment := models.Enroll{
			UserId:   &userId,
			CourseId: &courseId,
			CreatedAt: &time.Time{},
			UpdatedAt: &time.Time{},
		}

		if err := repo.db.Create(&enrollment).Error; err != nil {
			return err
		}

		return nil
	}

	// If the user is already enrolled, return an error
	return errors.New("user is already enrolled in this course")
}
