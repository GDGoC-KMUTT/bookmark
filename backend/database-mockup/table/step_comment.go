package table

import "time"

type StepComment struct {
	Id        *uint64    `gorm:"primaryKey"`
	StepId    *uint64    `gorm:"not null"`
	Step      *Step      `gorm:"foreignKey:StepId"`
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	Content   *string    `gorm:"type:TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
