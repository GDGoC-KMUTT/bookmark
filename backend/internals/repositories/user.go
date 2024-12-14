package repositories

import "backend/internals/db/models"

type UserRepository interface {
	Create(user *models.User) error
	First(user *models.User, conds ...interface{}) error
	//FindAll() (users []*entities.UserInfo, err error)
	//FindByEmail(email *string) (user *entities.UserInfo, err error)
	//FindById(id *primitive.ObjectID) (user *entities.UserInfo, err error)
	//DeleteById(id *primitive.ObjectID) (err error)
	//Update(user *entities.UserInfo) error
}
