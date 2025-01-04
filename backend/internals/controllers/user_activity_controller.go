package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/entities/payload"
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
		return &response.GenericError{
			Err:     err,
			Message: "Failed to fetch recent user activities",
		}
	}

	return response.Ok(c, recentActivities)
}

// CreateOrUpdateActivity
// @Summary Create or update a user activity
// @Description Create or update a user activity record based on stepId. User ID is extracted from the JWT.
// @Tags UserActivity
// @Accept json
// @Produce json
// @Param stepId path uint64 true "Step ID" example(123)
// @Success 200 {object} response.InfoResponse[string]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /user/activity/{stepId} [post]
func (c *UserActivityController) CreateOrUpdateActivity(ctx *fiber.Ctx) error {
	param := new(payload.UserActivityParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId parameter",
		}
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	err := c.userActivitySvc.UpdateUserActivity(uint64(userId), param.StepId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to update user activity",
		}
	}

	return response.Ok(ctx, "user activity updated successfully")
}
