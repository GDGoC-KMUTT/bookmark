package repositories

import "backend/internals/db/models"

// Renamed the interface to CoursePageRepositoryInterface to avoid conflict
type CoursePageRepositoryInterface interface {
	FindCoursePageInfoByCoursePageID(coursePageId string) (*models.Course, error)
	FindCoursePageContentByCoursePageID(coursePageId string) ([]models.CourseContent, error)
	FindSuggestCourseByFieldID(fieldId string) ([]models.Course, error)
}
