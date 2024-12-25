package models

import "time"

type Step struct {
	Id          *uint64    `gorm:"primaryKey"`
	ModuleId    *uint64    `gorm:"not null"`
	Module      *Module    `gorm:"foreignKey:ModuleId"`
	Title       *string    `gorm:"type:VARCHAR(255); not null"`
	Description *string    `gorm:"type:TEXT; null"`
	Content     *string    `gorm:"type:TEXT; null"` // Markdown
	Outcome     *string    `gorm:"type:TEXT; null"` // Markdown
	Check       *string    `gorm:"type:TEXT; null"` // Markdown
	Error       *string    `gorm:"type:TEXT; null"` // Markdown
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
}
