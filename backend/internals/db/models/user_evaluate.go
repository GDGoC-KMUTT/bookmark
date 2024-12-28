package models

import "time"

type UserEvaluate struct {
	Id             *uint64       `gorm:"primaryKey"`
	UserId         *uint64       `gorm:"index:idx_user_evaluate; not null"`
	User           *User         `gorm:"foreignKey:UserId"`
	StepEvaluateId *uint64       `gorm:"index:idx_user_evaluate; not null"`
	StepEvaluate   *StepEvaluate `gorm:"foreignKey:StepEvaluateId"`
	Content        *string       `gorm:"type:TEXT; not null"`
	Pass           *bool         `gorm:"null"`
	Comment        *string       `gorm:"type:TEXT; null"`
	CreatedAt      *time.Time    `gorm:"not null"`
	UpdatedAt      *time.Time    `gorm:"not null"`
}
