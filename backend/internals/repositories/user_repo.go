package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type userRepository struct {
	UserModel *gorm.DB
}

func NewUserRepository(userModel *gorm.DB) UserRepository {
	return &userRepository{
		UserModel: userModel,
	}
}

func (u *userRepository) Create(user *models.User) error {
	if err := u.UserModel.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) First(user *models.User, conds ...interface{}) error {
	if err := u.UserModel.First(user, conds...).Error; err != nil {
		return err
	}
	return nil
}
