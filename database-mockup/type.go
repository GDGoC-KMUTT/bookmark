package main

import "time"

type Article struct {
	Id        *uint64    `gorm:"primaryKey"`
	Title     *string    `gorm:"not null"`
	Href      *string    `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
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
type Course struct {
	Id        *uint64    `gorm:"primaryKey"`
	Name      *string    `gorm:"type:VARCHAR(255); not null"`
	FieldId   *uint64    `gorm:"not null"`
	Field     *Field     `gorm:"foreignKey:FieldId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
type Enroll struct {
	Id        *uint64    `gorm:"primaryKey"`
	UserId    *uint64    `gorm:"not null"`
	User      *User      `gorm:"foreignKey:UserId"`
	CourseId  *uint64    `gorm:"not null"`
	Course    *Course    `gorm:"foreignKey:CourseId"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
type Field struct {
	Id        *uint64    `gorm:"primaryKey"`
	Name      *string    `gorm:"type:VARCHAR(255); not null"`
	ImageUrl  *string    `gorm:"type:TEXT; null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
type Module struct {
	Id          *uint64    `gorm:"primaryKey"`
	Title       *string    `gorm:"type:VARCHAR(255); not null"`
	Description *string    `gorm:"type:TEXT; null"`
	ImageUrl    *string    `gorm:"type:TEXT; null"`
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
}
type StepAuthor struct {
	StepId *uint64 `gorm:"primaryKey"`
	Step   *Step   `gorm:"foreignKey:StepId"`
	UserId *uint64 `gorm:"primaryKey"`
	User   *User   `gorm:"foreignKey:UserId"`
}
type StepCommentUpvote struct {
	StepCommentId *uint64      `gorm:"primaryKey"`
	StepComment   *StepComment `gorm:"foreignKey:StepCommentId"`
	UserId        *uint64      `gorm:"primaryKey"`
	User          *User        `gorm:"foreignKey:UserId"`
	CreatedAt     *time.Time   `gorm:"not null"`
	UpdatedAt     *time.Time   `gorm:"not null"`
}
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

type StepEvaluate struct {
	Id        *uint64    `gorm:"primaryKey"`
	StepId    *uint64    `gorm:"index:idx_step_evaluate,unique; not null"`
	Step      *Step      `gorm:"foreignKey:StepId"`
	Order     *int       `gorm:"index:idx_step_evaluate,unique; not null"`
	Type      *string    `gorm:"type:VARCHAR(255) CHECK(type IN ('text', 'image')); not null"`
	Prompt    *string    `gorm:"type:TEXT; not null"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}
type Step struct {
	Id          *uint64    `gorm:"primaryKey"`
	ModuleId    *uint64    `gorm:"not null"`
	Module      *Module    `gorm:"foreignKey:ModuleId"`
	Title       *string    `gorm:"type:VARCHAR(255); not null"`
	Description *string    `gorm:"type:TEXT; null"`
	Gems        *int64     `gorm:"not null"`
	Content     *string    `gorm:"type:TEXT; null"` // Markdown
	Outcome     *string    `gorm:"type:TEXT; null"` // Markdown
	Check       *bool      `gorm:"not null"`        // Markdown
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
}
type UserEvaluate struct {
	UserId         *uint64       `gorm:"primaryKey"`
	User           *User         `gorm:"foreignKey:UserId"`
	StepEvaluateId *uint64       `gorm:"primaryKey"`
	StepEvaluate   *StepEvaluate `gorm:"foreignKey:StepEvaluateId"`
	Content        *string       `gorm:"type:TEXT; not null"`
	Check          *bool         `gorm:"not null"`
	Comment        *string       `gorm:"type:TEXT; null"`
	CreatedAt      *time.Time    `gorm:"not null"`
	UpdatedAt      *time.Time    `gorm:"not null"`
}
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
