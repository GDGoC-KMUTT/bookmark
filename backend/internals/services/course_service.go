package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"log"
)

type courseService struct {
	courseRepo repositories.CourseRepository
}

func NewCourseService(courseRepo repositories.CourseRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
	}
}

func (r *courseService) GetCourseByFieldId(fieldId *uint) ([]payload.CourseWithFieldType, error) {
	courses, fieldType, tx := r.courseRepo.FindCourseByFieldId(fieldId)
	if tx != nil {
		return nil, tx
	}
	var result []payload.CourseWithFieldType

	for _, course := range courses {
		result = append(result, payload.CourseWithFieldType{
			Id:         course.Id,
			Name:       course.Name,
			FieldId:    fieldType.Id,
			FieldName:  fieldType.Name,
			FieldImage: fieldType.ImageUrl,
		})
	}

	return result, nil
}

func (r *courseService) GetEnrollCourseByUserId(userId int) ([]*payload.Enroll, error) {
	log.Println("service userId: ", userId)
	enrollments, tx := r.courseRepo.FindEnrollCourseByUserId(userId) // Expecting a slice from the repository
	if tx != nil {
		// log.Println("Error fetching enrollment for userId:", userId)
		return nil, tx
	}
	if enrollments == nil {
        // log.Println("No enrollments found for userId:", userId)
        return nil, nil
    }

	var result []*payload.Enroll
	for _, enroll := range enrollments {
		result = append(result, &payload.Enroll{
			Id:       enroll.Id,
			UserId:   enroll.UserId,
			CourseId: enroll.CourseId,
		})
	}

	// log.Println("Enrollments found: ", result)
	return result, nil
}
