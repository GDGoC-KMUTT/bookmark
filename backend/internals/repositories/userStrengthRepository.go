package repositories

import (
	"backend/internals/db/models"
	"backend/internals/entities/response"
	"gorm.io/gorm"
	"log"
)

type UserStrengthRepository interface {
	GetStrengthDataByUserID(userId uint64) ([]response.StrengthDataResponse, error)
	GetSuggestionCourse(userId uint64) ([]models.Course, error)
}

type userStrengthRepository struct {
	db *gorm.DB
}

func NewUserStrengthRepository(db *gorm.DB) UserStrengthRepository {
	return &userStrengthRepository{
		db: db,
	}
}

// GetStrengthDataByUserID ดึงคะแนน strength ตามแต่ละ field type ที่ผู้ใช้ตอบถูก
func (r *userStrengthRepository) GetStrengthDataByUserID(userId uint64) ([]response.StrengthDataResponse, error) {
	// สร้างแผนที่เพื่อเก็บคะแนนตาม FieldType
	var strengthData []response.StrengthDataResponse

	var evaluations []struct {
		FieldName string
		TotalGems int64
	}

	err := r.db.
		Table("user_evaluates").
		Joins("JOIN step_evaluates ON step_evaluates.id = user_evaluates.step_evaluate_id").
		Joins("JOIN steps ON steps.id = step_evaluates.step_id").
		Joins("JOIN modules ON modules.id = steps.module_id").
		Joins("JOIN course_contents ON course_contents.module_id = modules.id").
		Joins("JOIN courses ON courses.id = course_contents.course_id").
		Joins("JOIN field_types ON field_types.id = courses.field_id").
		Where("user_evaluates.user_id = ? AND user_evaluates.pass = ?", userId, true).
		Select("field_types.name AS field_name, modules.title AS module_name, SUM(steps.gems) AS total_gems").
		Group("field_types.name, modules.title").
		Find(&evaluations).Error

	if err != nil {
		log.Printf("Error fetching evaluated steps for user %d: %v", userId, err)
		return nil, err
	}

	// สร้าง slice ของ StrengthDataDTO จากผลลัพธ์ที่ได้
	for _, evaluation := range evaluations {
		strengthData = append(strengthData, response.StrengthDataResponse{
			FieldName: evaluation.FieldName,
			TotalGems: evaluation.TotalGems,
		})
	}

	return strengthData, nil
}

func (r *userStrengthRepository) GetSuggestionCourse(userId uint64) ([]models.Course, error) {
	log.Printf("Fetching random course suggestions for user ID: %d", userId)

	var courses []models.Course
	err := r.db.
		Preload("Field").
		Order("RANDOM()").
		Limit(5).
		Find(&courses).Error

	if err != nil {
		log.Printf("Error fetching random courses for user %d: %v", userId, err)
		return nil, err
	}

	log.Printf("Found %d random courses for user %d", len(courses), userId)
	return courses, nil
}
