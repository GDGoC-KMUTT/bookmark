package repositories

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
)

type UserStrengthRepository interface {
	GetStrengthDataByUserID(userId uint64) ([]payload.StrengthFieldData, error)
	GetSuggestionCourse(userId uint64) ([]models.Course, error)
}
