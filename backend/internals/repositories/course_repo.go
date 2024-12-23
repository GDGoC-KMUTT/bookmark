package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
	// "log"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) FindCourseByFieldId(field_id *uint) ([]models.Course, *models.FieldType, error) {
	var courses []models.Course
	var fieldType *models.FieldType

	result := r.db.Find(&courses, "field_id = ?", field_id)
	if result.Error != nil {
		return nil, fieldType, result.Error
	}

	result = r.db.First(&fieldType, "id = ?", field_id)
	if result.Error != nil {
		return nil, fieldType, result.Error
	}

	return courses, fieldType, nil
}

func (r *courseRepository) FindEnrollCourseByUserId(userId int) ([]*models.Enroll, error) {
	var enrollments []*models.Enroll
	result := r.db.Where("user_id = ?", userId).Find(&enrollments)
	// log.Printf("Executing query to fetch all enrollments for userId: %d\n", userId)

	if result.RowsAffected == 0 {
		// log.Println("No records found for userId:", userId)
		return nil, nil
	}

	if result.Error != nil {
		// log.Println("Error fetching enroll records for userId:", userId, "Error:", result.Error)
		return nil, result.Error
	}

	// log.Printf("Found %d enrollments for userId: %d\n", result.RowsAffected, userId)
	return enrollments, nil
}

