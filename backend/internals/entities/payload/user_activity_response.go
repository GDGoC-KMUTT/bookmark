package payload

import "time"

type UserActivityResponse struct {
	CourseId    uint64     `json:"courseId"`
	ModuleId    uint64     `json:"moduleId"`
	StepID      uint64     `json:"stepId"`
	StepTitle   string     `json:"stepTitle"`
	ModuleTitle string     `json:"moduleTitle"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type UserActivitiesResponse struct {
	Activities []UserActivityResponse `json:"activities"`
}
