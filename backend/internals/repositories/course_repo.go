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

	if result.RowsAffected == 0 {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return enrollments, nil
}

func (r *courseRepository) FindCourseByCourseId(courseId *uint64) (*models.Course, error) {
	var course models.Course
	result := r.db.Where("id = ?", courseId).First(&course)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &course, nil
}

func (r *courseRepository) FindFieldByFieldId(fieldId uint64) (*models.FieldType, error) {
	var field models.FieldType
	result := r.db.Where("id = ?", fieldId).First(&field)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &field, nil
}
