package payload

type Course struct {
	Id      *uint64 `json:"id"`
	Name    *string `json:"name"`
	FieldId *uint64  `json:"fieldId"`
}

type CourseWithFieldType struct {
	Id            *uint64 `json:"id"`
	Name          *string `json:"name"`
	FieldId       *uint64 `json:"fieldId"`
	FieldName     *string `json:"fieldName"`
	FieldImageUrl *string `json:"fieldImageUrl"`
}

type TotalStepsByCourseIdPayload struct {
    CourseId   uint `json:"courseId"`
    TotalSteps int  `json:"total_steps"`
}

type CourseIdParam struct {
    CourseId uint `json:"courseId" param:"courseId"`
}
