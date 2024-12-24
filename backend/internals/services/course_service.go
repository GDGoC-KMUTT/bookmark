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

func (r *courseService) GetEnrollCourseByUserId(userId int) ([]*payload.EnrollwithCourse, error) {
	log.Println("service userId: ", userId)
	enrollments, tx := r.courseRepo.FindEnrollCourseByUserId(userId)
	if tx != nil {
		// log.Println("Error fetching enrollment for userId:", userId)
		return nil, tx
	}
	if enrollments == nil {
        // log.Println("No enrollments found for userId:", userId)
        return nil, nil
    }

	var result []*payload.EnrollwithCourse
	// for _, enroll := range enrollments {
	// 	result = append(result, &payload.Enroll{
	// 		Id:       enroll.Id,
	// 		UserId:   enroll.UserId,
	// 		CourseId: enroll.CourseId,
	// 	})
	// }
	for _, enroll := range enrollments {
		// Fetch course details (e.g., name) using courseId
		course, tx := r.courseRepo.FindCourseByCourseId(enroll.CourseId)
		if tx != nil {
			log.Println("Error fetching course details for courseId:", enroll.CourseId)
			return nil, tx
			// continue // Skip this enrollment if course details cannot be fetched
		}

		var fieldId *int64
		if course.FieldId != nil {
			// Explicit conversion from *uint64 to *int64
			fieldId = new(int64)
			*fieldId = int64(*course.FieldId)
		}

		var fieldImageURL *string
		var fieldName *string
		if fieldId != nil {
			// Assuming there's a method to get the field image URL by fieldId
			field, tx := r.courseRepo.FindFieldByFieldId(uint64(*fieldId))
			if tx != nil {
				log.Println("Error fetching field details for fieldId:", *fieldId)
				return nil, tx
			}
			// If field details are found, set the field image URL
			fieldImageURL = field.ImageUrl
			fieldName = field.Name
		}

		// Append enrollment data along with course name
		result = append(result, &payload.EnrollwithCourse{
			Id:       enroll.Id,
			UserId:   enroll.UserId,
			CourseId: enroll.CourseId,
			CourseName: &payload.Course{
				Id:   course.Id,
				Name: course.Name,
				FieldId: fieldId,
			},
			FieldImageURL: fieldImageURL,
			FieldName:     fieldName,
		})
	}

	// log.Println("Enrollments found: ", result)
	return result, nil
}
