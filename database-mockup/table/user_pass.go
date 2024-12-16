package table

import "time"

type UserPass struct {
	Id        *uint64    `gorm:"primaryKey"`
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId; not null"`
	Type      *string    `gorm:"type:VARCHAR(255) CHECK(type IN ('step', 'course', 'module')); not null"`
	Step      *Step      `gorm:"foreignKey:StepId; null"`
	StepId    *uint64    `gorm:"null"`
	Course    *Course    `gorm:"foreignKey:CourseId; null"`
	CourseId  *uint64    `gorm:"null"`
	Module    *Module    `gorm:"foreignKey:ModuleId; null"`
	ModuleId  *uint64    `gorm:"null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
