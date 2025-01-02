package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

type EnrollController struct {
	enrollService services.EnrollService
}

// NewEnrollController creates a new controller instance
func NewEnrollController(enrollService services.EnrollService) *EnrollController { // Accept the interface
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

// GetUserEnrollments
// @ID getUserEnrollments
// @Tags enroll
// @Summary GetUserEnrollments
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.EnrollmentListResponse]
// @Failure 400 {object} response.GenericError
// @Router /enrollments/enroll [get]
func (r *EnrollController) GetUserEnrollments(c *fiber.Ctx) error {
	// Get user information from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	// Fetch enrollments
	enrollments, err := r.enrollService.GetEnrollmentsByUserID(utils.Ptr(strconv.Itoa(int(userId))))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "Failed to fetch user enrollments",
		}
	}

	return response.Ok(c, enrollments)
}
