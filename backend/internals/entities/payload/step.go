package payload

type StepIdParam struct {
	StepId *uint64 `param:"stepId"`
}

type GetGemsResponse struct {
	TotalGems   *int `json:"totalGems"`
	CurrentGems *int `json:"currentGems"`
}
