package repositories

import (
	"backend/internals/db/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

type enrollRepository struct {
	db *gorm.DB
}

func NewEnrollRepository(db *gorm.DB) enrollRepository {
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

func (r *enrollRepository) FindEnrollmentsByUserID(userId *string) ([]models.Enroll, error) {
	var enrollments []models.Enroll
	err := r.db.
		Preload("Course").
		Where("user_id = ?", userId).
		Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollRepository) GetTotalStepsByCourseID(courseId uint64) (int64, error) {
	log.Printf("Fetching total steps for course ID: %d", courseId)
	var totalSteps int64
	err := r.db.
		Model(&models.Step{}).
		Where("module_id IN (SELECT module_id FROM course_contents WHERE course_id = ?)", courseId).
		Count(&totalSteps).Error
	if err != nil {
		log.Printf("Error fetching total steps for course %d: %v", courseId, err)
		return 0, err
	}
	log.Printf("Total steps for course %d: %d", courseId, totalSteps)
	return totalSteps, nil
}

func (r *enrollRepository) GetEvaluatedStepsByUserAndCourse(userId uint64, courseId uint64) (int64, error) {
	log.Printf("Fetching evaluated steps for user: %d, course: %d", userId, courseId)
	var evaluatedSteps int64
	err := r.db.
		Model(&models.UserEvaluate{}).
		Joins("JOIN step_evaluates ON step_evaluates.id = user_evaluates.step_evaluate_id").
		Joins("JOIN steps ON steps.id = step_evaluates.step_id").
		Joins("JOIN modules ON modules.id = steps.module_id").
		Joins("JOIN course_contents ON course_contents.module_id = modules.id").
		Where("user_evaluates.user_id = ? AND course_contents.course_id = ?", userId, courseId).
		Where("user_evaluates.pass = ?", true).
		Count(&evaluatedSteps).Error
	if err != nil {
		log.Printf("Error fetching evaluated steps for user %d, course %d: %v", userId, courseId, err)
		return 0, err
	}
	log.Printf("Evaluated steps for user %d, course %d: %d", userId, courseId, evaluatedSteps)
	return evaluatedSteps, nil
}

func (r *enrollRepository) FindEnrollmentsByUserID(userId *string) ([]models.Enroll, error) {
	var enrollments []models.Enroll
	err := r.db.
		Preload("Course").
		Where("user_id = ?", userId).
		Find(&enrollments).Error
	return enrollments, err
}
