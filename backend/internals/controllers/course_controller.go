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
// @Tags course
// @Summary Get all courses for a specific field_id
// @Accept json
// @Produce json
// @Param field_id path uint true "Field ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /course/field/{field_id} [get]
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

// GetAllFieldTypes
// @ID getAllFieldTypes
// @Tags course
// @Summary Get all field types
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.FieldType]
// @Failure 400 {object} response.GenericError
// @Router /course/field_types [get]
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
