package models

import "time"

type UserEvaluate struct {
	UserId         *uint64       `gorm:"primaryKey"`
	User           *User         `gorm:"foreignKey:UserId"`
	StepEvaluateId *uint64       `gorm:"primaryKey"`
	StepEvaluate   *StepEvaluate `gorm:"foreignKey:StepEvaluateId"`
	Content        *string       `gorm:"type:TEXT; not null"`
	Pass           *bool         `gorm:"null"`
	Comment        *string       `gorm:"type:TEXT; null"`
	CreatedAt      *time.Time    `gorm:"not null"`
	UpdatedAt      *time.Time    `gorm:"not null"`
}
