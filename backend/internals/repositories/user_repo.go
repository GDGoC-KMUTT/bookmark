package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindUserByID(id *string) (*models.User, error) {
	user := new(models.User)
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *userRepository) FindFirstUserByOid(oid *string) (*models.User, error) {
	user := new(models.User)
	result := r.db.First(&user, "oid = ?", *oid)
	return user, result.Error
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetTotalGemsByUserID(userID uint) (uint64, error) {
	var totalGems uint64
	err := r.db.Table("user_passes").
		Joins("INNER JOIN step_evaluates ON user_passes.step_id = step_evaluates.step_id").
		Where("user_passes.user_id = ?", userID).
		Select("COALESCE(SUM(step_evaluates.gem), 0) AS total_gems"). // Handle NULL values with COALESCE
		Scan(&totalGems).Error
	if err != nil {
		return 0, err
	}
	return totalGems, nil
}

func (r *userRepository) GetUserCompletedSteps(userID uint) ([]models.UserActivity, error) {
	var userActivities []models.UserActivity
	result := r.db.Where("user_id = ?", userID).Find(&userActivities)
	if result.Error != nil {
		return nil, result.Error
	}

	return userActivities, nil
}

func (r *userRepository) GetUserPassByUserID(userID uint, stepId uint) (int64, error) {
	var count int64
	err := r.db.Table("user_passes").
		Joins("INNER JOIN step_evaluates ON user_passes.step_id = step_evaluates.step_id").
		Where("user_passes.user_id = ? and user_passes.step_id = ?", userID, stepId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
