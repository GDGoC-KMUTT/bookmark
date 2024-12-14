package models

import "time"

type Course struct {
	Id        *uint64    `gorm:"primaryKey"`
	Name      *string    `gorm:"type:VARCHAR(255); not null"`
	FieldId   *uint64    `gorm:"not null"`
	Field     *FieldType `gorm:"foreignKey:FieldId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
