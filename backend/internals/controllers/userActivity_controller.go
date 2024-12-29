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
// @Summary Get the most recent activities for a user
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.UserActivitiesResponse]
// @Failure 400 {object} response.GenericError
// @Router /user/recent-activities [get]
func (r *UserActivityController) GetRecentActivity(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	recentActivities, err := r.userActivitySvc.GetRecentActivitiesByUserID(utils.Ptr(strconv.Itoa(int(userId))))
	if err != nil {
		return c.JSON(response.GenericError{
			Err:     err,
			Message: "Failed to fetch recent user activities",
		})
	}

	return response.Ok(c, recentActivities)
}
