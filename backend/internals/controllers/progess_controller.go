package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ProgressController struct {
	progressService services.ProgressService
}

func NewProgressController(progressService services.ProgressService) *ProgressController {
	return &ProgressController{
		progressService: progressService,
	}
}

// GetCompletionPercentage
// @ID getCompletionPercentage
// @Tags progress
// @Summary Get course completion percentage
// @Description Calculates the course completion percentage for a user based on completed steps.
// @Accept json
// @Produce json
// @Param courseID path int true "Course ID"
// @Success 200 {object} map[string]float64 "completion_percentage"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /progress/{courseID}/percentage [get]
func (pc *ProgressController) GetCompletionPercentage(c *fiber.Ctx) error {
	// Extract user from JWT token using c.Locals
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["userId"].(float64))

	// Extract courseID from URL parameter
	courseIDParam := c.Params("courseID")
	courseID, err := strconv.ParseUint(courseIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid courseID",
		})
	}

	percentage, err := pc.progressService.GetCompletionPercentage(userId, uint(courseID))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Code:    "PROGRESS_FETCH_FAILED",
			Message: "failed to get completion percentage",
		}
	}

	// Return the percentage as a response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"completion_percentage": percentage,
	})
}
