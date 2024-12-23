package payload

type ProgressResponse struct {
    CompletionPercentage float64 `json:"completion_percentage"`
    TotalSteps           int     `json:"total_steps"`
    CompletedSteps       int     `json:"completed_steps"`
}