package services

type EnrollServices interface {
	EnrollUser(userId uint64, courseId uint64) error
}
