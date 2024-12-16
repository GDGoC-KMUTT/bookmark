package table

import "time"

type Enroll struct {
	Id        *uint64    `gorm:"primaryKey"`
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	CourseId  *uint64    `gorm:"not null"`
	Course    *Course    `gorm:"foreignKey:CourseId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
