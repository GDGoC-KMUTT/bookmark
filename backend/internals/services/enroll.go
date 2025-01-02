package services

import "backend/internals/entities/payload"

type EnrollService interface {
	GetEnrollmentsByUserID(userId *string) (payload.EnrollmentListResponse, error)
	EnrollUser(userId uint, courseId uint64) error
}
