package models

import "time"

type Article struct {
	Id        *uint64    `gorm:"primaryKey"`
	Title     *string    `gorm:"not null"`
	Href      *string    `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
