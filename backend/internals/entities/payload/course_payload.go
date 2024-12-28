package payload

type Course struct {
	Id      uint64  `json:"id"`
	Name    string  `json:"name"`
	FieldId uint64  `json:"field_id"`
	Field   *string `json:"field"`
}

type CourseContent struct {
	CourseId uint64  `json:"course_id"`
	Order    int64   `json:"order"`
	Type     string  `json:"type"`
	Text     *string `json:"text,omitempty"`   // Nullable for text type
	ModuleId *uint64 `json:"module_id,omitempty"` // Nullable for module id
}
