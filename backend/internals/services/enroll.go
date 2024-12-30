package services

type EnrollServices interface {
	EnrollUser(userId uint, courseId uint64) error
}
