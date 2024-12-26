package payload

type Course struct {
	Id      *uint64 `json:"id"`
	Name    *string `json:"name"`
	FieldId *uint64  `json:"field_id"`
}

type CourseWithFieldType struct {
	Id         *uint64 `json:"id"`
	Name       *string `json:"name"`
	FieldId    *uint64 `json:"field_id"`
	FieldName  *string `json:"field_name"`
	FieldImage *string `json:"field_image"`
}

type TotalStepsByCourseIdPayload struct {
    CourseId   uint `json:"courseId"`
    TotalSteps int  `json:"total_steps"`
}

type CourseIdParam struct {
    CourseId uint `json:"courseId" param:"courseId"`
}

type Enroll struct {
	Id         *uint64 `json:"id"`
	UserId     *uint64 `json:"userId"`
	CourseId   *uint64 `json:"courseId"`
}

type EnrollwithCourse struct {
	Id         	  *uint64 `json:"id"`
	UserId     	  *uint64 `json:"userId"`
	CourseId   	  *uint64 `json:"courseId"`
	CourseName 	  *Course `json:"courseName"`
	FieldName  	  *string `json:"fieldName"`
	FieldImageURL *string `json:"fieldImageUrl"`
}