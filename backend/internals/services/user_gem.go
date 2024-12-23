package services

import "backend/internals/entities/payload"

type GemService interface {
	GetTotalGems(userID uint) (*payload.GemTotal, error)
}
