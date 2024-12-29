package controllers

import (
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type EnrollController struct {
	enrollService *services.EnrollService
}

// NewCourseEnrollmentController creates a new controller instance
func NewEnrollController(enrollService *services.EnrollService) *EnrollController {
	return &EnrollController{enrollService: enrollService}
}

// @Summary Enroll a user in a course
// @Description Enroll a user in a specified course by user ID and course ID
// @Tags Enroll
// @Accept json
// @Produce json
// @Param userId path uint64 true "User ID" example(123)
// @Param courseId path uint64 true "Course ID" example(456)
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/enroll/{userId}/{courseId} [post]
func (c *EnrollController) EnrollInCourse(ctx *fiber.Ctx) error {
	userId, err := strconv.ParseUint(ctx.Params("userId"), 10, 64)
	if err != nil || userId <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	courseId, err := strconv.ParseUint(ctx.Params("courseId"), 10, 64)
	if err != nil || courseId <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid course ID"})
	}

	// Call the EnrollUser service method
	err = c.enrollService.EnrollUser(userId, courseId)
	if err != nil {
		// Handle user already enrolled scenario (409 Conflict)
		if err.Error() == "user is already enrolled in this course" {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already enrolled"})
		}
		// Handle other errors (internal server error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user enrolled successfully"})
}
