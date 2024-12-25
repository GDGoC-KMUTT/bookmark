package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
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

func (r *courseService) GetCurrentCourse(userID uint) (*payload.Course, error) {
	course, err := r.courseRepo.GetCurrentCourse(userID)
	if err != nil {
		return nil, err
	}

	courseDetails := &payload.Course{
		Id:   course.Id,
		Name: course.Name,
	}

	return courseDetails, nil
}

func (r *courseService) GetTotalStepsByCourseId(courseID uint) (*payload.TotalStepsByCourseIdPayload, error) {
	totalSteps, err := r.courseRepo.GetTotalStepsByCourseId(courseID)
	if err != nil {
		return nil, err
	}

	return &payload.TotalStepsByCourseIdPayload{
		CourseId:   courseID,
		TotalSteps: totalSteps,
	}, nil
}

func (r *courseService) GetEnrollCourseByUserId(userId int) ([]*payload.EnrollwithCourse, error) {
	enrollments, tx := r.courseRepo.FindEnrollCourseByUserId(userId)
	if tx != nil {
		// log.Println("Error fetching enrollments for userId:", userId, "Error:", tx)
		return nil, tx
	}
	if enrollments == nil {
        return []*payload.EnrollwithCourse{}, nil
    }

	var result []*payload.EnrollwithCourse
	for _, enroll := range enrollments {
		course, tx := r.courseRepo.FindCourseByCourseId(enroll.CourseId)
		if tx != nil {
			// log.Println("Error fetching course details for courseId:", enroll.CourseId)
			return nil, tx
		}

		var fieldId *int64
		if course.FieldId != nil {
			fieldId = new(int64)
			*fieldId = int64(*course.FieldId)
		}

		var fieldImageURL *string
		var fieldName *string
		if fieldId != nil {
			field, tx := r.courseRepo.FindFieldByFieldId(uint64(*fieldId))
			if tx != nil {
				// log.Println("Error fetching field details for fieldId:", *fieldId)
				return nil, tx
			}
			if field != nil {
                fieldImageURL = field.ImageUrl
                fieldName = field.Name
            }
		}

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

	return result, nil
}
