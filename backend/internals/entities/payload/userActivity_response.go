package payload

type UserActivityResponse struct {
	StepID      uint64 `json:"stepId"`
	StepTitle   string `json:"stepTitle"`
	ModuleTitle string `json:"moduleTitle"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type UserActivitiesResponse struct {
	Activities []UserActivityResponse `json:"activities"`
}
