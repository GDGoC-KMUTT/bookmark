package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
)

// CoursePageController handles course page-related endpoints
type CoursePageController struct {
	coursePageSvc *services.CoursePageService
}

// NewCoursePageController initializes a new CoursePageController
func NewCoursePageController(coursePageSvc *services.CoursePageService) *CoursePageController {
	return &CoursePageController{
		coursePageSvc: coursePageSvc,
	}
}

// GetCoursePageInfo
// @ID getCoursePageInfo
// @Tags course_page
// @Summary Get course page information
// @Accept json
// @Produce json
// @Param coursePageId path string true "Course Page ID"
// @Success 200 {object} response.InfoResponse[payload.CoursePage]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /courses/{coursePageId}/info [get]
func (c *CoursePageController) GetCoursePageInfo(ctx *fiber.Ctx) error {
	coursePageId := ctx.Params("coursePageId")

	// Call service to get course page info
	coursePageInfo, err := c.coursePageSvc.GetCoursePageInfo(coursePageId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get course page info",
		})
	}

	// Dereference coursePageInfo and pass as value to InfoResponse
	return ctx.JSON(&response.InfoResponse[payload.CoursePage]{
		Data: *coursePageInfo, // Dereference the pointer here
	})
}

// GetCoursePageContent
// @ID getCoursePageContent
// @Tags course_page
// @Summary Get course page content
// @Accept json
// @Produce json
// @Param coursePageId path string true "Course Page ID"
// @Success 200 {object} response.InfoResponse[[]payload.CoursePageContent]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /courses/{coursePageId}/content [get]
func (c *CoursePageController) GetCoursePageContent(ctx *fiber.Ctx) error {
	coursePageId := ctx.Params("coursePageId")

	// Call service to get course page content
	coursePageContent, err := c.coursePageSvc.GetCoursePageContent(coursePageId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get course page content",
		})
	}

	return ctx.JSON(&response.InfoResponse[[]payload.CoursePageContent]{
		Data: coursePageContent,
	})
}
