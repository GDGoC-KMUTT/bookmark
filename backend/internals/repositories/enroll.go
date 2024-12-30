package repositories

type EnrollRepo interface {
	EnrollUser(userId uint, courseId uint64) error
}
