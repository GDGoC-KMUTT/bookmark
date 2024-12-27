package payload

type Enrollment struct {
	Id        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	CourseID  uint64 `json:"course_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
