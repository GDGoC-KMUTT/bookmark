package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
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
		return nil, nil, result.Error
	}

	result = r.db.First(&fieldType, "id = ?", field_id)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return courses, fieldType, nil
}
