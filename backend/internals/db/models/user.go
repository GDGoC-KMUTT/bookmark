package models

import "time"

type User struct {
	Id        *uint64    `gorm:"primaryKey"`
	Oid       *string    `gorm:"type:VARCHAR(255); index:idx_user_oid,unique; not null"` // OAuth ID from Microsoft Sign-in
	Firstname *string    `gorm:"type:VARCHAR(255); not null"`
	Lastname  *string    `gorm:"type:VARCHAR(255); not null"`
	Email     *string    `gorm:"type:VARCHAR(255); index:idx_user_email,unique; not null"`
	PhotoUrl  *string    `gorm:"type:TEXT; null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
