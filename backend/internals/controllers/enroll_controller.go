package controllers

import (
	"backend/internals/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type EnrollController struct {
	enrollService services.EnrollServices // Use the interface here
}

// NewEnrollController creates a new controller instance
func NewEnrollController(enrollService services.EnrollServices) *EnrollController { // Accept the interface
	return &EnrollController{enrollService: enrollService}
}

// @Summary Enroll a user in a course
// @Description Enroll a user in a specified course by course ID. User ID is extracted from the JWT.
// @Tags Enroll
// @Accept json
// @Produce json
// @Param courseId path uint64 true "Course ID" example(456)
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /enroll/{courseId} [post]
func (c *EnrollController) EnrollInCourse(ctx *fiber.Ctx) error {
	// Extract user from context
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	userId, ok := claims["userId"].(float64)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	// Parse courseId from the request path
	courseId, err := strconv.ParseUint(ctx.Params("courseId"), 10, 64)
	if err != nil || courseId <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid course ID"})
	}

	// Call the EnrollUser service method
	err = c.enrollService.EnrollUser(uint(userId), courseId)
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
