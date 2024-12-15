package models

import "time"

type UserActivity struct {
	UserId    *uint64    `gorm:"index:idx_user_activity_user_id"`
	User      *User      `gorm:"foreignKey:UserId"`
	StepId    *uint64    `gorm:"not null"`
	Step      *Step      `gorm:"foreignKey:StepId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
