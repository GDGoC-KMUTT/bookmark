package repositories

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"gorm.io/gorm"
	"log"
)

type userStrengthRepository struct {
	db *gorm.DB
}

func NewUserStrengthRepository(db *gorm.DB) UserStrengthRepository {
	return &userStrengthRepository{
		db: db,
	}
}

// GetStrengthDataByUserID ดึงคะแนน strength ตามแต่ละ field type ที่ผู้ใช้ตอบถูก
func (r *userStrengthRepository) GetStrengthDataByUserID(userId uint64) ([]payload.StrengthFieldData, error) {
	var strengthData []payload.StrengthFieldData

	var evaluations []struct {
		FieldName string
		TotalGems int64
	}

	err := r.db.Table("user_passes").
		Joins("INNER JOIN step_evaluates ON user_passes.step_id = step_evaluates.step_id").
		Joins("INNER JOIN steps ON user_passes.step_id = steps.id").
		Joins("INNER JOIN modules ON steps.module_id = modules.id").
		Joins("INNER JOIN course_contents ON course_contents.module_id = modules.id").
		Joins("INNER JOIN courses ON courses.id = course_contents.course_id").
		Joins("INNER JOIN field_types ON field_types.id = courses.field_id").
		Where("user_passes.user_id = ?", userId).
		Select("field_types.name, COALESCE(SUM(step_evaluates.gem), 0) AS total_gems").
		Group("field_types.name").
		Find(&evaluations).Error

	if err != nil {
		log.Printf("Error fetching evaluated steps for user %d: %v", userId, err)
		return nil, err
	}

	// If no evaluations found, fetch all default field types
	if len(evaluations) == 0 {
		var fieldTypes []models.FieldType
		err := r.db.Find(&fieldTypes).Error
		if err != nil {
			log.Printf("Error fetching default field types: %v", err)
			return nil, err
		}

		strengthData = make([]payload.StrengthFieldData, len(fieldTypes))
		for i, fieldType := range fieldTypes {
			strengthData[i] = payload.StrengthFieldData{
				FieldName: *fieldType.Name,
				TotalGems: 0,
			}
		}

		log.Printf("No passed evaluations found for user %d. Returning %d default field types.", userId, len(strengthData))
		return strengthData, nil
	}

	// Convert evaluations to response format
	strengthData = make([]payload.StrengthFieldData, 0, len(evaluations))
	for _, evaluation := range evaluations {
		strengthData = append(strengthData, payload.StrengthFieldData{
			FieldName: evaluation.FieldName,
			TotalGems: evaluation.TotalGems,
		})
	}

	return strengthData, nil
}

// GetSuggestionCourse returns random course suggestions
func (r *userStrengthRepository) GetSuggestionCourse(userId uint64) ([]models.Course, error) {
	log.Printf("Fetching random course suggestions for user ID: %d", userId)

	var courses []models.Course
	err := r.db.
		Preload("Field"). // Preload the field relationship
		Order("RANDOM()"). // Random order
		Limit(5). // Limit to 5 courses
		Find(&courses).Error

	if err != nil {
		log.Printf("Error fetching random courses for user %d: %v", userId, err)
		return nil, err
	}

	log.Printf("Found %d random courses for user %d", len(courses), userId)
	return courses, nil
}
