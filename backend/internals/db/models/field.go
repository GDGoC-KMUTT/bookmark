package models

import "time"

type Field struct {
	Id        *uint64    `gorm:"primaryKey"`
	Name      *string    `gorm:"type:VARCHAR(255); not null"`
	ImageUrl  *string    `gorm:"type:TEXT; null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
