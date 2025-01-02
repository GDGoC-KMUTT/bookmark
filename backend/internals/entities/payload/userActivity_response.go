package payload

import "time"

type UserActivityResponse struct {
	StepID      uint64     `json:"stepId"`
	StepTitle   string     `json:"stepTitle"`
	ModuleTitle string     `json:"moduleTitle"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type UserActivitiesResponse struct {
	Activities []UserActivityResponse `json:"activities"`
}
