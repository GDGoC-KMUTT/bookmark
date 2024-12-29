package payload

type CourseResponse struct {
	ID    uint64        `json:"id"`
	Name  string        `json:"name"`
	Field FieldResponse `json:"field"`
}

type FieldResponse struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	ImageUrl *string `json:"imageUrl,omitempty"`
}

type CourseSuggestionResponse struct {
	Courses    []CourseResponse `json:"courses"`
	TotalCount int64            `json:"totalCount"`
}
