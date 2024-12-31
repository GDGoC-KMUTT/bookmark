package payload

type EnrollRequest struct {
    UserId   uint64 `json:"userId"`
    CourseId uint64 `json:"courseId"`
}
