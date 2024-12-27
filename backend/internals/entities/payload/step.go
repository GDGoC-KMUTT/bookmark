package payload

type StepIdParam struct {
	StepId *uint64 `param:"stepId"`
}

type GetGemsResponse struct {
	TotalGems   *int `json:"totalGems"`
	CurrentGems *int `json:"currentGems"`
}

type StepCommentInfo struct {
	StepCommentId *uint64      `json:"stepCommentId"`
	UserInfo      *CommentedBy `json:"userInfo"`
	Comment       *string      `json:"comment"`
	UpVote        *int         `json:"upVote"`
}

type CommentedBy struct {
	UserId    *uint64 `json:"userId"`
	FirstName *string `json:"firstName"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	PhotoUrl  *string `json:"photoUrl"`
}

type Comment struct {
	StepId  *uint64 `json:"stepId" validate:"required"`
	Content *string `json:"content" validate:"required"`
}

type UpVoteComment struct {
	StepCommentId *uint64 `json:"stepCommentId" validate:"required"`
}
