package payload

type ModuleResponse struct {
	Id          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	ImageUrl    *string `json:"image_url,omitempty"`
}
