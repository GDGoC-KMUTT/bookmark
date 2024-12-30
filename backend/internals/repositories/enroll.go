package repositories

type EnrollRepo interface {
	EnrollUser(userId, courseId uint64) error
}
