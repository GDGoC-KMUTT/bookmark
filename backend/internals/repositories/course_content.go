package repositories

type CourseContentRepository interface {
	GetCourseIdByModuleId(moduleId *uint64) (*uint64, error)
}
