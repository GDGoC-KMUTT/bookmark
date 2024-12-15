package repositories

import "backend/internals/db/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByID(id *string) (*models.User, error)
	FindFirstUserByOid(oid *string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}
