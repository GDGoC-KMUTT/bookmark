package db

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

var UserModel *gorm.DB

func AssignModel() {
	UserModel = Gorm.Model(new(models.User))
}
