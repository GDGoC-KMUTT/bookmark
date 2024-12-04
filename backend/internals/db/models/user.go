package models

type User struct {
	UserId    *string `json:"userId" gorm:"not null;primaryKey"`
	FirstName *string `json:"firstName" gorm:"not null"`
	LastName  *string `json:"lastName" gorm:"not null"`
	Email     *string `json:"email" gorm:"not null"`
	Avatar    *string `json:"avatar" gorm:""`
}
