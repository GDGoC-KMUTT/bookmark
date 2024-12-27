package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

type UserActivityController struct {
	userActivitySvc services.UserActivityService
}

func NewUserActivityController(userActivitySvc services.UserActivityService) *UserActivityController {
	return &UserActivityController{
		userActivitySvc: userActivitySvc,
	}
}

// GetRecentActivity
// @ID getRecentActivity
// @Tags user-activity
// @Summary Get the most recent activity for a user
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[response.UserActivityResponse]
// @Failure 400 {object} response.GenericError
// @Router /user/{userId}/recent-activity [get]
func (r *UserActivityController) GetRecentActivity(c *fiber.Ctx) error {
	// Get user information from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	// Fetch the most recent user activity
	recentActivity, err := r.userActivitySvc.GetRecentActivityByUserID(utils.Ptr(strconv.Itoa(int(userId))))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "Failed to fetch recent user activity",
		}
	}

	return response.Ok(c, recentActivity)
}
