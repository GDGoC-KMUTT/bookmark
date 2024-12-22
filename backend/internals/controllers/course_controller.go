package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"

	"github.com/gofiber/fiber/v2"
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
// @Tags courses
// @Summary Get all courses for a specific field_id
// @Accept json
// @Produce json
// @Param field_id path uint true "Field ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldImage]
// @Failure 400 {object} response.GenericError
// @Router /courses/field/{field_id} [get]
func (r *CourseController) GetCoursesByFieldId(c *fiber.Ctx) error {

	param := new(payload.FieldIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid field_id parameter",
		}
	}

	courses, err := r.courseSvc.GetCourseByFieldId(param.FieldId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get courses",
		}
	}

	return response.Ok(c, courses)
}
