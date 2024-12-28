package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
	"fmt"
)

// Merge the courseService and CourseService structs
type CourseService struct {
	courseRepo    repositories.CourseRepository
	fieldTypeRepo repositories.FieldTypeRepository // Added this field from the second file
}

// Merge the two NewCourseService functions to account for the fieldTypeRepo parameter
func NewCourseService(courseRepo repositories.CourseRepository, fieldTypeRepo repositories.FieldTypeRepository) *CourseService {
	return &CourseService{
		courseRepo:    courseRepo,
		fieldTypeRepo: fieldTypeRepo,
	}
}

// From the first file: GetCourseInfo
func (s *CourseService) GetCourseInfo(courseId string) (*payload.Course, error) {
	// Fetch course from the repository
	courseEntity, err := s.courseRepo.FindCourseInfoByCourseID(courseId)
	if err != nil {
		return nil, err
	}

	// Map to payload.Course
	return &payload.Course{
		Id:      *courseEntity.Id,
		Name:    *courseEntity.Name,
		FieldId: *courseEntity.FieldId,
		Field:   courseEntity.Field.Name, // Assuming Field has a `Name` field
	}, nil
}

// From the first file: GetCourseContent
func (s *CourseService) GetCourseContent(courseId string) ([]payload.CourseContent, error) {
	// Fetch course content from the repository
	contentEntities, err := s.courseRepo.FindCourseContentByCourseID(courseId)
	if err != nil {
		return nil, err
	}

	// Map to []payload.CourseContent
	contents := make([]payload.CourseContent, 0, len(contentEntities))
	for _, content := range contentEntities {
		contents = append(contents, payload.CourseContent{
			CourseId: *content.CourseId,
			Order:    *content.Order,
			Type:     *content.Type,
			Text:     content.Text,
			ModuleId: content.ModuleId,
		})
	}

	return contents, nil
}

// From the second file: GetCoursesByFieldId
func (s *CourseService) GetCoursesByFieldId(fieldId uint) ([]payload.CourseWithFieldType, error) {
	courses, fieldType, tx := s.courseRepo.FindCourseByFieldId(fieldId)
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

// From the second file: GetAllFieldTypes
func (s *CourseService) GetAllFieldTypes() ([]payload.FieldType, error) {
	fieldTypes, tx := s.fieldTypeRepo.FindAllFieldTypes()
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

// From the second file: GetCurrentCourse
func (s *CourseService) GetCurrentCourse(userID uint) (*payload.Course, error) {
	course, err := s.courseRepo.GetCurrentCourse(userID)
	if err != nil {
		return nil, err
	}

	courseDetails := &payload.Course{
		Id:   course.Id,
		Name: course.Name,
	}

	return courseDetails, nil
}

// From the second file: GetTotalStepsByCourseId
func (s *CourseService) GetTotalStepsByCourseId(courseID uint) (*payload.TotalStepsByCourseIdPayload, error) {
	totalSteps, err := s.courseRepo.GetTotalStepsByCourseId(courseID)
	if err != nil {
		return nil, err
	}

	return &payload.TotalStepsByCourseIdPayload{
		CourseId:   courseID,
		TotalSteps: totalSteps,
	}, nil
}

// From the second file: GetEnrollCourseByUserId
func (s *CourseService) GetEnrollCourseByUserId(userId int) ([]*payload.EnrollwithCourse, error) {
	enrollments, tx := s.courseRepo.FindEnrollCourseByUserId(userId)
	if tx != nil {
		fmt.Printf("Error fetching enrollments for userId: %d, Error: %v\n", userId, tx)
		return nil, tx
	}
	if enrollments == nil {
		return []*payload.EnrollwithCourse{}, nil
	}

	var result []*payload.EnrollwithCourse
	for _, enroll := range enrollments {
		course, tx := s.courseRepo.FindCourseByCourseId(enroll.CourseId)
		if tx != nil {
			fmt.Printf("error fetching course details, Error: %v\n", tx)
			return nil, tx
		}

		if course.FieldId == nil {
			return nil, tx
		}

		field, tx := s.courseRepo.FindFieldByFieldId(course.FieldId)
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
