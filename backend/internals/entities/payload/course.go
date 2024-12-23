package payload

type Course struct {
	Id      *uint64 `json:"id"`
	Name    *string `json:"name"`
	FieldId *int64  `json:"field_id"`
}

type CourseWithFieldType struct {
	Id         *uint64 `json:"id"`
	Name       *string `json:"name"`
	FieldId    *uint64 `json:"field_id"`
	FieldName  *string `json:"field_name"`
	FieldImage *string `json:"field_image"`
}

