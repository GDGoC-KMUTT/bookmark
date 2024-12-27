package response

type UserActivityResponse struct {
	StepID      uint64 `json:"step_id"`
	StepTitle   string `json:"step_title"`
	ModuleTitle string `json:"module_title"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
