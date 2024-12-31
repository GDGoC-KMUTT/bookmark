package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
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

// EnrollInCourse
// @Summary Enroll a user in a course
// @Description Enroll a user in a specified course by course ID. User ID is extracted from the JWT.
// @Tags Enroll
// @Accept json
// @Produce json
// @Param courseId path uint64 true "Course ID" example(456)
// @Success 200 {object} response.InfoResponse[string]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /enroll/{courseId} [post]
func (c *EnrollController) EnrollInCourse(ctx *fiber.Ctx) error {
	param := new(payload.CourseIdParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid courseId parameter",
		}
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	// Call the EnrollUser service method
	err := c.enrollService.EnrollUser(uint(userId), uint64(param.CourseId))
	if err != nil {
		// Handle user already enrolled scenario (409 Conflict)
		if err.Error() == "user is already enrolled in this course" {
			return &response.GenericError{
				Err:     err,
				Message: "user already enrolled",
			}
		}
		// Handle other errors (internal server error)
		return &response.GenericError{
			Err:     err,
			Message: "failed to enroll course",
		}
	}

	return response.Ok(ctx, "user enrolled successfully")
}
