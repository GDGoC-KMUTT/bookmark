package repositories

import (
	"backend/internals/db/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type enrollRepository struct {
	db *gorm.DB
}

func NewEnrollRepository(db *gorm.DB) EnrollRepo {
	return &enrollRepository{
		db: db,
	}
}

// EnrollUser enrolls a user in a course
func (repo *enrollRepository) EnrollUser(userId uint, courseId uint64) error {
	// Check if the user is already enrolled
	var existingEnrollment models.Enroll
	result := repo.db.Where("user_id = ? AND course_id = ?", userId, courseId).First(&existingEnrollment)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// If no enrollment found, create a new enrollment record
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		enrollment := models.Enroll{
			UserId:    uint64Ptr(uint64(userId)),
			CourseId:  &courseId,
			CreatedAt: timePtr(time.Now()),
			UpdatedAt: timePtr(time.Now()),
		}

		if err := repo.db.Create(&enrollment).Error; err != nil {
			return err
		}

		return nil
	}

	// If the user is already enrolled, return an error
	return errors.New("user is already enrolled in this course")
}

// Helper function to create a pointer to a uint64 value
func uint64Ptr(v uint64) *uint64 {
	return &v
}

// Helper function to create a pointer to a time.Time value
func timePtr(t time.Time) *time.Time {
	return &t
}
