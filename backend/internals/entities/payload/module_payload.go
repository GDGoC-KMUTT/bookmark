package payload

type ModuleResponse struct {
	Id          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	ImageUrl    *string `json:"imageUrl,omitempty"`
}

type ModuleIdParam struct {
	ModuleId *uint64 `param:"moduleId" json:"moduleId"`
}
