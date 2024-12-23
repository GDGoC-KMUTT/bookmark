package repositories

import (
	"gorm.io/gorm"
)

type gemRepository struct {
	db *gorm.DB
}

func NewGemRepository(db *gorm.DB) GemRepository {
	return &gemRepository{
		db: db,
	}
}

// GetTotalGemsByUserID calculates the total gems for a given user
func (r *gemRepository) GetTotalGemsByUserID(userID uint) (uint64, error) {
	var totalGems uint64
	// Use COALESCE to convert NULL to 0
	err := r.db.Table("user_passes").
		Joins("INNER JOIN steps ON user_passes.step_id = steps.id").
		Where("user_passes.user_id = ?", userID).
		Select("COALESCE(SUM(steps.gems), 0) AS total_gems"). // Convert NULL to 0
		Scan(&totalGems).Error
	if err != nil {
		return 0, err
	}
	return totalGems, nil
}
