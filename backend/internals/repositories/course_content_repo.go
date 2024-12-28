package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type courseContentRepo struct {
	db *gorm.DB
}

func NewCourseContentRepository(db *gorm.DB) CourseContentRepository {
	return &courseContentRepo{
		db: db,
	}
}

func (r *courseContentRepo) GetCourseIdByModuleId(moduleId *uint64) (*uint64, error) {
	courseContent := new(models.CourseContent)
	result := r.db.First(&courseContent, "module_id = ? ", moduleId)
	return courseContent.CourseId, result.Error
}
