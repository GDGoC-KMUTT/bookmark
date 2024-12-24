package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
)

// CourseController handles course-related endpoints
type CourseController struct {
	courseSvc *services.CourseService
}

// NewCourseController initializes a new CourseController
func NewCourseController(courseSvc *services.CourseService) *CourseController {
	return &CourseController{
		courseSvc: courseSvc,
	}
}

// GetCourseInfo
// @ID getCourseInfo
// @Tags course
// @Summary Get course information
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} response.InfoResponse[payload.Course]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /course/{courseId}/info [get]
func (c *CourseController) GetCourseInfo(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	// Call service to get course info
	courseInfo, err := c.courseSvc.GetCourseInfo(courseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get course info",
		})
	}

	// Dereference courseInfo and pass as value to InfoResponse
	return ctx.JSON(&response.InfoResponse[payload.Course]{
		Data: *courseInfo,  // Dereference the pointer here
	})
}


// GetCourseContent
// @ID getCourseContent
// @Tags course
// @Summary Get course content
// @Accept json
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseContent]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /course/{courseId}/content [get]
func (c *CourseController) GetCourseContent(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	// Call service to get course content
	courseContent, err := c.courseSvc.GetCourseContent(courseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get course content",
		})
	}

	return ctx.JSON(&response.InfoResponse[[]payload.CourseContent]{
		Data: courseContent,
	})
}
