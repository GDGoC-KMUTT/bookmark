package repositories

import "backend/internals/db/models"

type CoursePageRepo interface {
	FindCoursePageInfoByCoursePageID(coursePageId string) (*models.Course, error)
	FindCoursePageContentByCoursePageID(coursePageId string) ([]models.CourseContent, error)
	FindSuggestCourseByFieldID(fieldId string) ([]models.Course, error)
}
