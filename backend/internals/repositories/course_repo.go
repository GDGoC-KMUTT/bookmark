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

func (r *courseRepository) GetCurrentCourse(userID uint) (*models.Course, error) {
	var userActivity models.UserActivity
	result := r.db.Where("user_id = ?", userID).Order("updated_at desc").First(&userActivity)
	if result.Error != nil {
		return nil, result.Error
	}

	var step models.Step
	result = r.db.First(&step, "id = ?", userActivity.StepId) 
	if result.Error != nil {
		return nil, result.Error
	}

	var module models.Module
	result = r.db.First(&module, "id = ?", step.ModuleId)
	if result.Error != nil {
		return nil, result.Error
	}

	var courseContent models.CourseContent
	result = r.db.First(&courseContent, "module_id = ?", module.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	var course models.Course
	result = r.db.First(&course, "id = ?", courseContent.CourseId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &course, nil
}

func (r *courseRepository) GetAllCourseSteps(courseID uint) ([]models.Step, error) {
    var courseContent models.CourseContent
    result := r.db.First(&courseContent, "course_id = ?", courseID)
    if result.Error != nil {
		return []models.Step{}, result.Error
    }

    var steps []models.Step
    result = r.db.Where("module_id = ?", courseContent.ModuleId).Find(&steps)
    if result.Error != nil {
        return []models.Step{}, result.Error
    }

    return steps, nil
}

func (r *courseRepository) GetTotalStepsByCourseId(courseID uint) (int, error) {
    steps, err := r.GetAllCourseSteps(courseID + 1)
    if err != nil {
        return 0, err
    }

    return len(steps), nil
}