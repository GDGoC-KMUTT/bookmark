package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

// Constructor for the repository
func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// FindCourseInfoByCourseID - fetches course details by course ID
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

// FindCourseContentByCourseID - fetches course content by course ID
func (r *CourseRepository) FindCourseContentByCourseID(courseId string) ([]models.CourseContent, error) {
	var contents []models.CourseContent
	err := r.db.Raw(`SELECT * FROM course_contents WHERE course_id = ? ORDER BY "order" ASC`, courseId).Scan(&contents).Error

	if err != nil {
		return nil, err
	}
	return contents, nil
}

// FindCourseByFieldId - fetches all courses by a specific field ID
func (r *CourseRepository) FindCourseByFieldId(fieldId uint) ([]models.Course, *models.FieldType, error) {
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

// GetCurrentCourse - fetches the current course a user is enrolled in
func (r *CourseRepository) GetCurrentCourse(userID uint) (*models.Course, error) {
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

// GetAllCourseSteps - fetches all steps of a course by course ID
func (r *CourseRepository) GetAllCourseSteps(courseID uint) ([]models.Step, error) {
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

// GetTotalStepsByCourseId - fetches the total number of steps in a course
func (r *CourseRepository) GetTotalStepsByCourseId(courseID uint) (int, error) {
	steps, err := r.GetAllCourseSteps(courseID)
	if err != nil {
		return 0, err
	}

	return len(steps), nil
}

// FindEnrollCourseByUserId - fetches all enrollments for a user
func (r *CourseRepository) FindEnrollCourseByUserId(userId int) ([]*models.Enroll, error) {
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

// FindCourseByCourseId - fetches a specific course by course ID
func (r *CourseRepository) FindCourseByCourseId(courseId *uint64) (*models.Course, error) {
	var course models.Course
	result := r.db.Where("id = ?", courseId).First(&course)

	if result.Error != nil {
		return nil, result.Error
	}

	return &course, nil
}

// FindFieldByFieldId - fetches the field type by field ID
func (r *CourseRepository) FindFieldByFieldId(fieldId *uint64) (*models.FieldType, error) {
	var field models.FieldType
	result := r.db.Where("id = ?", fieldId).First(&field)

	if result.Error != nil {
		return nil, result.Error
	}

	return &field, nil
}
