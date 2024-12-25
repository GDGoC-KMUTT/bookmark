package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type userEvaluateRepo struct {
	db *gorm.DB
}

func NewUserEvaluateRepo(db *gorm.DB) UserEvaluateRepository {
	return &userEvaluateRepo{
		db: db,
	}
}

func (r *userEvaluateRepo) GetUserEvalByUserId(userId *uint64) ([]*models.UserEvaluate, error) {
	userEvals := make([]*models.UserEvaluate, 0)

	result := r.db.Find(&userEvals, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return userEvals, nil
}
