package payload

type EnrollmentListDTO struct {
	Id         uint64  `json:"id"`
	UserID     uint64  `json:"userId"`
	CourseID   uint64  `json:"courseId"`
	CourseName string  `json:"courseName"`
	Progress   float64 `json:"progress"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type EnrollmentListResponse struct {
	Enrollments []EnrollmentListDTO `json:"enrollments"`
}
