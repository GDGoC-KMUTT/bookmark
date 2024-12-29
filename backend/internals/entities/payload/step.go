package payload

type StepIdParam struct {
	StepId *uint64 `param:"stepId"`
}

type UserEvalIdsBody struct {
	UserEvalIds *string `json:"userEvalIds" validate:"required"`
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
	HasUpVoted    *bool        `json:"hasUpVoted"`
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

type StepInfo struct {
	Step       *StepDetail `json:"step"`
	Authors    []*UserInfo `json:"authors"`
	UserPassed []*UserInfo `json:"userPassed"`
}

type StepDetail struct {
	StepId      *uint64 `json:"stepId"`
	ModuleId    *uint64 `json:"moduleId"`
	Banner      *string `json:"banner"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Content     *string `json:"content"`
	Outcome     *string `json:"outcome"`
	Check       *string `json:"check"`
	Error       *string `json:"error"`
}

type StepInfoQuery struct {
	CourseId *uint64 `query:"courseId" validate:"required"`
	ModuleId *uint64 `query:"moduleId" validate:"required"`
}

type SubmitStepEval struct {
	StepId     *uint64 `json:"stepId" validate:"required"`
	StepEvalId *uint64 `json:"stepEvalId" validate:"required"`
	Content    *string `json:"content"`
}

type StepEvalInfo struct {
	StepEvalId  *uint64         `json:"stepEvalId"`
	StepId      *uint64         `json:"stepId"`
	Order       *int            `json:"order"`
	Question    *string         `json:"question"`
	Type        *string         `json:"type"`
	Instruction *string         `json:"instruction"`
	UserEval    *UserEvalResult `json:"userEval"`
}

type UserEvalResult struct {
	Content *string `json:"content"`
	Pass    *bool   `json:"pass"`
	Comment *string `json:"comment "`
}

type CreateUserEvalReq struct {
	UserId     *float64 `json:"userId"`
	StepEvalId *uint64  `json:"stepEvalId"`
	Content    *string  `json:"content"`
}

type CreateUserEvalRes struct {
	UserEvalId *uint64 `json:"userEvalId"`
}

type UserInfo struct {
	UserId    *uint64 `json:"userId"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	PhotoUrl  *string `json:"photoUrl"`
}
