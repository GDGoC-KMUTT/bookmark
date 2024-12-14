package models

import "time"

type Module struct {
	Id          *uint64    `gorm:"primaryKey"`
	Title       *string    `gorm:"type:VARCHAR(255); not null"`
	Description *string    `gorm:"type:TEXT; null"`
	ImageUrl    *string    `gorm:"type:TEXT; null"`
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
}
