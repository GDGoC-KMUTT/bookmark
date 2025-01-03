package services

import "backend/internals/entities/payload"

type UserStrengthService interface {
	GetStrengthDataByUserID(userId uint64) (*payload.StrengthDataResponse, error)
	GetSuggestionCourse(userId uint64) ([]payload.CourseResponse, error)
}
