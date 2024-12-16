package table

import "time"

type CourseContent struct {
	CourseId  *uint64    `gorm:"primaryKey"`
	Course    *Course    `gorm:"foreignKey:CourseId"`
	Order     *int64     `gorm:"primaryKey"`
	Type      *string    `gorm:"type:VARCHAR(255) CHECK(type IN ('text', 'module')); not null"`
	Text      *string    `gorm:"type:TEXT; null"`
	ModuleId  *uint64    `gorm:"null"`
	Module    *Module    `gorm:"foreignKey:ModuleId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
