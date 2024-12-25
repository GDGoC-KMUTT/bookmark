package services

type StepService interface {
	GetGems(stepId *uint64, userId *float64) (*int, *int, error)
}
