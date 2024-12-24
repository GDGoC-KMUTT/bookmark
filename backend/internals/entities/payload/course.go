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

type Enroll struct {
	Id         *uint64 `json:"id"`
	UserId     *uint64 `json:"user_id"`
	CourseId   *uint64 `json:"course_id"`
}

type EnrollwithCourse struct {
	Id         	  *uint64 `json:"id"`
	UserId     	  *uint64 `json:"user_id"`
	CourseId   	  *uint64 `json:"course_id"`
	CourseName 	  *Course `json:"course_name"`
	FieldName  	  *string `json:"field_name"`
	FieldImageURL *string `json:"field_image_url"`
}