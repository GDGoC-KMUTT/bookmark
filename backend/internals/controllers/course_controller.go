package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CourseController struct {
	courseSvc services.CourseService
}

func NewCourseController(courseSvc services.CourseService) CourseController {
	return CourseController{
		courseSvc: courseSvc,
	}
}

// GetCoursesByFieldId
// @ID getCoursesByFieldId
// @Tags course
// @Summary Get all courses for a specific fieldId
// @Accept json
// @Produce json
// @Param fieldId path uint true "Field ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /courses/field/{fieldId} [get]
func (r *CourseController) GetCoursesByFieldId(c *fiber.Ctx) error {
	param := new(payload.FieldIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid fieldId parameter",
		}
	}

	courses, err := r.courseSvc.GetCoursesByFieldId(uint(param.FieldId))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get course by fieldId",
		}
		// return c.Status(http.StatusInternalServerError).JSON(response.GenericError{
		// 	Err:     err,
		// 	Message: "failed to get course by fieldId",
		// })
	}

	return response.Ok(c, courses)
}

// GetAllFieldTypes
// @ID getAllFieldTypes
// @Tags course
// @Summary Get all field types
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.FieldType]
// @Failure 400 {object} response.GenericError
// @Router /courses/field_types [get]
func (r *CourseController) GetAllFieldTypes(c *fiber.Ctx) error {
	fieldTypes, err := r.courseSvc.GetAllFieldTypes()
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get all field types",
		}
	}
	return response.Ok(c, fieldTypes)
}

// GetCurrentCourse
// @ID getCurrentCourse
// @Tags courses
// @Summary Get the current course the user is in based on their latest activity
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[payload.Course]
// @Failure 400 {object} response.GenericError
// @Router /courses/current [get]
func (r *CourseController) GetCurrentCourse(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["userId"].(float64)

	course, err := r.courseSvc.GetCurrentCourse(uint(userID))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to fetch current course",
		}
	}

	return response.Ok(c, course)
}

// GetTotalStepsByCourseId
// @ID getTotalStepsByCourseId
// @Tags courses
// @Summary Get the total steps for a specific course by course_id
// @Accept json
// @Produce json
// @Param course_id path uint true "Course ID"
// @Success 200 {object} response.InfoResponse[payload.TotalStepsByCourseIdPayload]
// @Failure 400 {object} response.GenericError
// @Router /courses/{courseId}/total-steps [get]
func (r *CourseController) GetTotalStepsByCourseId(c *fiber.Ctx) error {
	param := new(payload.CourseIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid course_id parameter",
		}
	}

	totalSteps, err := r.courseSvc.GetTotalStepsByCourseId(param.CourseId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to fetch total steps",
		}
	}

	return response.Ok(c, totalSteps)
}
