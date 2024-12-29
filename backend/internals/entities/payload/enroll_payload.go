// package payload

// // EnrollRequest is used for enrolling a user in a course
// type EnrollRequest struct {
//     UserId   uint64 `json:"userId" binding:"required"`
//     CourseId uint64 `json:"courseId" binding:"required"`
// }

package payload

type EnrollRequest struct {
    UserId   uint64 `json:"userId"`
    CourseId uint64 `json:"courseId"`
}
