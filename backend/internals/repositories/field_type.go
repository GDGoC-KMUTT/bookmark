package repositories

import "backend/internals/db/models"


type FieldTypeRepository interface {
	FindAllFieldTypes() ([]models.FieldType, error)
}
