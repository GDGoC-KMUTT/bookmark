package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
    db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
    return CourseRepository{
        db: db,
    }
}

func (r *CourseRepository) FindCourseInfoByCourseID(courseId string) (*models.Course, error) {
	var course models.Course
	err := r.db.Preload("Field"). // Preload the associated FieldType
		Where("id = ?", courseId).
		First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}


func (r *CourseRepository) FindCourseContentByCourseID(courseId string) ([]models.CourseContent, error) {
	var contents []models.CourseContent
	err := r.db.Raw(`SELECT * FROM course_contents WHERE course_id = ? ORDER BY "order" ASC`, courseId).Scan(&contents).Error

	if err != nil {
		return nil, err
	}
	return contents, nil
}



