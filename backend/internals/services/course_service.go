package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"fmt"
)

type courseService struct {
	courseRepo    repositories.CourseRepository
	fieldTypeRepo repositories.FieldTypeRepository
}

func NewCourseService(courseRepo repositories.CourseRepository, fieldTypeRepo repositories.FieldTypeRepository) CourseService {
	return &courseService{
		courseRepo:    courseRepo,
		fieldTypeRepo: fieldTypeRepo,
	}
}

func (r *courseService) GetCoursesByFieldId(fieldId uint) ([]payload.CourseWithFieldType, error) {
	courses, fieldType, tx := r.courseRepo.FindCourseByFieldId(fieldId)
	if tx != nil {
		return nil, tx
	}
	var result []payload.CourseWithFieldType

	for _, course := range courses {
		result = append(result, payload.CourseWithFieldType{
			Id:            course.Id,
			Name:          course.Name,
			FieldId:       fieldType.Id,
			FieldName:     fieldType.Name,
			FieldImageUrl: fieldType.ImageUrl,
		})
	}

	return result, nil
}

func (r *courseService) GetAllFieldTypes() ([]payload.FieldType, error) {
	fieldTypes, tx := r.fieldTypeRepo.FindAllFieldTypes()
	if tx != nil {
		return nil, tx
	}

	var result []payload.FieldType
	for _, fieldType := range fieldTypes {
		result = append(result, payload.FieldType{
			Id:       fieldType.Id,
			Name:     fieldType.Name,
			ImageUrl: fieldType.ImageUrl,
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
		fmt.Printf("Error fetching enrollments for userId: %d, Error: %v\n", userId, tx)
		return nil, tx
	}
	if enrollments == nil {
		return []*payload.EnrollwithCourse{}, nil
	}

	var result []*payload.EnrollwithCourse
	for _, enroll := range enrollments {
		course, tx := r.courseRepo.FindCourseByCourseId(enroll.CourseId)
		if tx != nil {
			fmt.Printf("error fetching course details, Error: %v\n", tx)
			return nil, tx
		}

		if course.FieldId == nil {
			return nil, tx
		}

		field, tx := r.courseRepo.FindFieldByFieldId(course.FieldId)
		if tx != nil {
			fmt.Printf("error fetching field details, Error: %v\n", tx)
			return nil, tx
		}

		if field == nil {
			return nil, tx
		}

		result = append(result, &payload.EnrollwithCourse{
			Id:       enroll.Id,
			UserId:   enroll.UserId,
			CourseId: enroll.CourseId,
			CourseName: &payload.Course{
				Id:   course.Id,
				Name: course.Name,
				FieldId: course.FieldId,
			},
			FieldImageURL: field.ImageUrl,
			FieldName:     field.Name,
		})
	}

	return result, nil
}
