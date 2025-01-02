package services

type UserActivityService interface {
	UpdateUserActivity(userId uint64, stepId uint64) error
}
