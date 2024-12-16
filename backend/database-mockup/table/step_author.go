package table

type StepAuthor struct {
	StepId *uint64 `gorm:"primaryKey"`
	Step   *Step   `gorm:"foreignKey:StepId"`
	UserId *uint64 `gorm:"primaryKey"`
	User   *User   `gorm:"foreignKey:UserId"`
}
