package table

import "time"

type StepEvaluate struct {
	Id        *uint64    `gorm:"primaryKey"`
	StepId    *uint64    `gorm:"index:idx_step_evaluate,unique; not null"`
	Step      *Step      `gorm:"foreignKey:StepId"`
	Order     *int       `gorm:"index:idx_step_evaluate,unique; not null"`
	Type      *string    `gorm:"type:VARCHAR(255) CHECK(type IN ('text', 'image')); not null"`
	Prompt    *string    `gorm:"type:TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
