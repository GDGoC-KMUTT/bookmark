package repositories

type GemRepository interface {
	GetTotalGemsByUserID(userID uint) (uint64, error)
}
