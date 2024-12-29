package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"fmt"
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

// GetSuggestCoursesByFieldId
// @ID getSuggestCoursesByFieldId
// @Tags course_page
// @Summary Get suggest courses by field ID
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.SuggestCourse]
// @Failure 400 {object} response.GenericError
// @Router /courses/suggest/{fieldId} [get]
func (r *CoursePageController) GetSuggestCoursesByFieldId(c *fiber.Ctx) error {
    // Get fieldId as a string from the URL path parameter
    fieldIdStr := c.Params("fieldId")

    // Check if the fieldId parameter is provided
    if fieldIdStr == "" {
        return c.Status(fiber.StatusBadRequest).JSON(response.GenericError{
            Err:     fmt.Errorf("missing fieldId parameter"),
            Message: "fieldId parameter is required",
        })
    }

    // Optional: If the service expects fieldId as a string, no need for conversion.
    // Call the service to get the suggested courses by fieldId
    suggestInfo, err := r.coursePageSvc.GetSuggestCourseByFieldId(fieldIdStr) // Pass fieldId as a string
    if err != nil {
        fmt.Printf("Failed to fetch suggest courses for fieldId %v: %v\n", fieldIdStr, err)
        return c.Status(fiber.StatusInternalServerError).JSON(response.GenericError{
            Err:     err,
            Message: "Failed to fetch suggest course",
        })
    }

    // Return successful response
    return c.Status(fiber.StatusOK).JSON(response.InfoResponse[[]payload.SuggestCourse]{
        Data: suggestInfo,
    })
}
