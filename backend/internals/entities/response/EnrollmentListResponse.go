package response

type EnrollmentListDTO struct {
	Id         uint64  `json:"id"`
	UserID     uint64  `json:"user_id"`
	CourseID   uint64  `json:"course_id"`
	CourseName string  `json:"course_name"`
	Progress   float64 `json:"progress"` // New field for progress
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type EnrollmentListResponse struct {
	Enrollments []EnrollmentListDTO `json:"enrollments"`
}
