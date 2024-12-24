package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
)

// Controller
type CourseController struct {
    courseSvc *services.CourseService // Use a pointer here
}

func NewCourseController(courseSvc *services.CourseService) CourseController {
    return CourseController{
        courseSvc: courseSvc,
    }
}


// GetCourseInfo
// @Router /course/:courseId/info [get]
func (c *CourseController) GetCourseInfo(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	// Call service to get course info
	courseInfo, err := c.courseSvc.GetCourseInfo(courseId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get course info",
		}
	}

	return response.Ok(ctx, courseInfo)
}

// GetCourseContent
// @Router /course/:courseId/content [get]
func (c *CourseController) GetCourseContent(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")

	// Call service to get course content
	courseContent, err := c.courseSvc.GetCourseContent(courseId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get course content",
		}
	}

	return response.Ok(ctx, courseContent)
}
