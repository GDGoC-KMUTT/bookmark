package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

type EnrollController struct {
	enrollSvc services.EnrollService
}

func NewEnrollController(enrollSvc services.EnrollService) EnrollController {
	return EnrollController{
		enrollSvc: enrollSvc,
	}
}

// GetUserEnrollments
// @ID getUserEnrollments
// @Tags enroll
// @Summary GetUserEnrollments
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.Enrollment]
// @Failure 400 {object} response.GenericError
// @Router /enrollments/ [get]
func (r *EnrollController) GetUserEnrollments(c *fiber.Ctx) error {
	// Get user information from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	// Fetch enrollments
	enrollments, err := r.enrollSvc.GetEnrollmentsByUserID(utils.Ptr(strconv.Itoa(int(userId))))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "Failed to fetch user enrollments",
		}
	}
	
	return response.Ok(c, enrollments)
}
