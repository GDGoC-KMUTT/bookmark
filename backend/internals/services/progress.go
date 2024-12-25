package services

type ProgressService interface {
	GetCompletionPercentage(userID uint, courseID uint) (float64, error)
}
