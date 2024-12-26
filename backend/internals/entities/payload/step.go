package payload

type StepIdParam struct {
	StepId *uint64 `param:"stepId"`
}

type GetGemsResponse struct {
	TotalGems   *int `json:"totalGems"`
	CurrentGems *int `json:"currentGems"`
}

type StepCommentInfo struct {
	UserInfo *CommentedBy `json:"userInfo"`
	Comment  *string      `json:"comment"`
}

type CommentedBy struct {
	UserId    *uint64 `json:"userId"`
	FirstName *string `json:"firstName"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	PhotoUrl  *string `json:"photoUrl"`
}
