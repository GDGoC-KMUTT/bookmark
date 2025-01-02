package repositories

type UserActivityRepository interface {
	UpdateUserActivity(userId uint64, stepId uint64) error
}
