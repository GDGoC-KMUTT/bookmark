package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type CoursePageRepository struct {
	db *gorm.DB
}

func NewCoursePageRepository(db *gorm.DB) CoursePageRepository {
	return CoursePageRepository{
		db: db,
	}
}

func (r *CoursePageRepository) FindCoursePageInfoByCoursePageID(coursePageId string) (*models.Course, error) {
	var coursePage models.Course
	err := r.db.Preload("Field"). // Preload the associated FieldType
		Where("id = ?", coursePageId).
		First(&coursePage).Error
	if err != nil {
		return nil, err
	}
	return &coursePage, nil
}

func (r *CoursePageRepository) FindCoursePageContentByCoursePageID(coursePageId string) ([]models.CourseContent, error) {
	var contents []models.CourseContent
	err := r.db.Raw(`SELECT * FROM course_contents WHERE course_id = ? ORDER BY "order" ASC`, coursePageId).Scan(&contents).Error

	if err != nil {
		return nil, err
	}
	return contents, nil
}
