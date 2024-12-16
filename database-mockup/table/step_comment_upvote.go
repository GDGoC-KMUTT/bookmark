package table

import "time"

type StepCommentUpvote struct {
	StepCommentId *uint64      `gorm:"primaryKey"`
	StepComment   *StepComment `gorm:"foreignKey:StepCommentId"`
	UserId        *uint64      `gorm:"primaryKey"`
	User          *User        `gorm:"foreignKey:UserId"`
	CreatedAt     *time.Time   `gorm:"not null"`
	UpdatedAt     *time.Time   `gorm:"not null"`
}
