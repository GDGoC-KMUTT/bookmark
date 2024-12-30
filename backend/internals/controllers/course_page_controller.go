package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"fmt"
	"strings"
	"strconv"
)

// CoursePageController handles course page-related endpoints
type CoursePageController struct {
	coursePageSvc services.CoursePageServices
}

func NewCoursePageController(service services.CoursePageServices) *CoursePageController {
	return &CoursePageController{
		coursePageSvc: service, // Assign the service parameter to the field
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

	// Check if the coursePageId is numeric
	if _, err := strconv.Atoi(coursePageId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&response.GenericError{
			Message: fmt.Sprintf("course page with ID %s not found", coursePageId),
		})
	}

	// Validate the Authorization token
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&response.GenericError{
			Message: "missing or invalid authorization token",
		})
	}

	// Call the service to get course page info
	coursePageInfo, err := c.coursePageSvc.GetCoursePageInfo(coursePageId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(fiber.StatusNotFound).JSON(&response.GenericError{
				Err:     err,
				Message: fmt.Sprintf("course page with ID %s not found", coursePageId),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get course page info",
		})
	}

	return ctx.JSON(&response.InfoResponse[payload.CoursePage]{
		Data: *coursePageInfo,
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

	coursePageContent, err := c.coursePageSvc.GetCoursePageContent(coursePageId)
	if err != nil {
		// Handle other errors (e.g., internal server errors)
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get course page content",
		})
	}

	// Handle empty content case
	if len(coursePageContent) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(&response.GenericError{
			Message: fmt.Sprintf("no content found for course page ID %s", coursePageId),
		})
	}

	return ctx.JSON(&response.InfoResponse[[]payload.CoursePageContent]{
		Data: coursePageContent,
	})
}


// GetSuggestCoursesByFieldId
// @ID getSuggestCoursesByFieldId
// @Tags course_page
// @Summary Get suggest courses by field ID
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.SuggestCourse]
// @Failure 400 {object} response.GenericError
// @Router /courses/suggest/{fieldId} [get]
func (c *CoursePageController) GetSuggestCoursesByFieldId(ctx *fiber.Ctx) error {
	// Validate the Authorization token
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&response.GenericError{
			Message: "missing or invalid authorization token",
		})
	}

	fieldIdStr := ctx.Params("fieldId")

	// Validate if fieldId is numeric
	if _, err := strconv.Atoi(fieldIdStr); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&response.GenericError{
			Message: fmt.Sprintf("fieldId %s is invalid", fieldIdStr),
		})
	}

	suggestCourses, err := c.coursePageSvc.GetSuggestCourseByFieldId(fieldIdStr)
	if err != nil {
		// Handle "not found" errors
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(fiber.StatusNotFound).JSON(&response.GenericError{
				Err:     err,
				Message: fmt.Sprintf("no suggestions found for field ID %s", fieldIdStr),
			})
		}
		// Handle other errors
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to fetch suggested courses",
		})
	}

	return ctx.JSON(&response.InfoResponse[[]payload.SuggestCourse]{
		Data: suggestCourses,
	})
}
