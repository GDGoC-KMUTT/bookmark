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
		Joins("INNER JOIN steps ON user_passes.step_id = steps.id").
		Where("user_passes.user_id = ?", userID).
		Select("COALESCE(SUM(steps.gems), 0) AS total_gems"). // Convert NULL to 0
		Scan(&totalGems).Error
	if err != nil {
		return 0, err
	}
	return totalGems, nil
}

