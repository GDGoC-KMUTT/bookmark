package models

import "time"

type StepEvaluate struct {
	Id          *uint64    `gorm:"primaryKey"`
	StepId      *uint64    `gorm:"index:idx_step_evaluate; not null"`
	Step        *Step      `gorm:"foreignKey:StepId"`
	Gem         *int       `gorm:"not null"`
	Order       *int       `gorm:"index:idx_step_evaluate; not null"`
	Question    *string    `gorm:"type:TEXT; not null"`
	Type        *string    `gorm:"type:VARCHAR(255) CHECK(type IN ('check', 'text', 'image')); not null"`
	Instruction *string    `gorm:"type:TEXT; not null"`
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
}
