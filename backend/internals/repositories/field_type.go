package repositories

import "backend/internals/db/models"


type FieldTypeRepository interface {
	FindAllFieldType() ([]models.FieldType, error)
}
