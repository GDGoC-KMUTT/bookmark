package repositories

import (
	"backend/internals/db/models"

	"gorm.io/gorm"
)

type fieldTypeRepository struct {
	db *gorm.DB
}

func NewFieldTypeRepository(db *gorm.DB) FieldTypeRepository {
	return &fieldTypeRepository{
		db: db,
	}
}

func (r *fieldTypeRepository) FindAllFieldTypes() ([]models.FieldType, error) {
	var fieldTypes []models.FieldType

	result := r.db.Find(&fieldTypes)
	if result.Error != nil {
		return nil, result.Error
	}

	// Use a map to remove duplicates
	uniqueFieldTypesMap := make(map[string]models.FieldType)
	for _, fieldType := range fieldTypes {
		uniqueFieldTypesMap[*fieldType.Name] = fieldType
	}
	// Convert map back to a slice
	var uniqueFieldTypes []models.FieldType
	for _, fieldType := range uniqueFieldTypesMap {
		uniqueFieldTypes = append(uniqueFieldTypes, fieldType)
	}
	return uniqueFieldTypes, nil
}
