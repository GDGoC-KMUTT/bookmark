package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
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
	param := new(payload.CourseIdParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid courseId parameter",
		}
	}

	// Call the service to get course page info
	coursePageInfo, err := c.coursePageSvc.GetCoursePageInfo(strconv.Itoa(int(param.CourseId)))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &response.GenericError{
				Err:     err,
				Message: fmt.Sprintf("course page with ID %d not found", param.CourseId),
			}
		}
		return &response.GenericError{
			Err:     err,
			Message: "failed to get course page info",
		}
	}

	return response.Ok(ctx, coursePageInfo)
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
	param := new(payload.CourseIdParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid courseId parameter",
		}
	}

	coursePageContent, err := c.coursePageSvc.GetCoursePageContent(strconv.Itoa(int(param.CourseId)))
	if err != nil {
		// Handle other errors (e.g., internal server errors)
		return &response.GenericError{
			Err:     err,
			Message: "failed to get course page content",
		}
	}

	// Handle empty content case
	if len(coursePageContent) == 0 {
		return &response.GenericError{
			Message: fmt.Sprintf("no content found for course page ID %d", param.CourseId),
		}
	}

	return response.Ok(ctx, coursePageContent)
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
	param := new(payload.FieldIdParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid fieldId parameter",
		}
	}

	suggestCourses, err := c.coursePageSvc.GetSuggestCourseByFieldId(strconv.Itoa(int(param.FieldId)))
	if err != nil {
		// Handle "not found" errors
		if strings.Contains(err.Error(), "not found") {
			return &response.GenericError{
				Err:     err,
				Message: fmt.Sprintf("no suggestions found for field ID %d", param.FieldId),
			}
		}
		// Handle other errors
		return &response.GenericError{
			Err:     err,
			Message: "failed to fetch suggested courses",
		}
	}

	return response.Ok(ctx, suggestCourses)
}
