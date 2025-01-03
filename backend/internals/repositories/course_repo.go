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

func (r *courseRepository) FindCourseByFieldId(fieldId uint) ([]models.Course, *models.FieldType, error) {
	var courses []models.Course
	var fieldType *models.FieldType

	result := r.db.Find(&courses, "field_id = ?", fieldId)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	result = r.db.First(&fieldType, "id = ?", fieldId)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return courses, fieldType, nil
}

func (r *courseRepository) GetCurrentCourse(userID uint) (*models.Course, error) {
	var userActivity models.UserActivity
	result := r.db.Where("user_id = ?", userID).Order("updated_at desc").First(&userActivity)
	if result.Error != nil {
		return nil, nil
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
	var courseContents []models.CourseContent
	result := r.db.Where("course_id = ?", courseID).Find(&courseContents)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(courseContents) == 0 {
		return nil, nil
	}

	var moduleIDs []uint64
	for _, content := range courseContents {
		if content.ModuleId != nil {
			moduleIDs = append(moduleIDs, *content.ModuleId)
		}
	}

	var steps []models.Step
	result = r.db.Where("module_id IN ?", moduleIDs).Find(&steps)
	if result.Error != nil {
		return nil, result.Error
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
		return nil, result.Error
	}

	return &course, nil
}

func (r *courseRepository) FindFieldByFieldId(fieldId *uint64) (*models.FieldType, error) {
	var field models.FieldType
	result := r.db.Where("id = ?", fieldId).First(&field)

	if result.Error != nil {
		return nil, result.Error
	}

	return &field, nil
}

func (r *courseRepository) FindCoursesByFieldId(fieldId uint64) ([]models.Course, error) {
    var courses []models.Course
    result := r.db.Where("field_id = ?", fieldId).Find(&courses)
    if result.Error != nil {
        return nil, result.Error
    }
    return courses, nil
}

