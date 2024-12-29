package payload

type CoursePage struct {
	Id      uint64  `json:"id"`
	Name    string  `json:"name"`
	FieldId uint64  `json:"fieldId"`
	Field   *string `json:"field"`
}

type CoursePageContent struct {
	CoursePageId uint64  `json:"coursePageId"`
	Order        int64   `json:"order"`
	Type         string  `json:"type"`
	Text         *string `json:"text,omitempty"`   // Nullable for text type
	ModuleId     *uint64 `json:"moduleId,omitempty"` // Nullable for module id
}

type SuggestCourse struct {
    Id           uint64   `json:"id"`
    Name         string   `json:"name"`
    FieldId      uint64   `json:"fieldId"`
    FieldName    *string  `json:"fieldName"`
    FieldImageUrl *string `json:"fieldImageUrl"`
}
