package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserActivityController struct {
	userActivityService services.UserActivityService 
}

func NewUserActivityController(userActivityService services.UserActivityService) *UserActivityController {
	return &UserActivityController{userActivityService: userActivityService}
}

// UpdateUserActivity
// @Summary Create or update a user activity
// @Description Update a user activity record based on stepId. User ID is extracted from the JWT.
// @Tags UserActivity
// @Accept json
// @Produce json
// @Param stepId path uint64 true "Step ID" example(123)
// @Success 200 {object} response.InfoResponse[string]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /userActivity/{stepId} [post]
func (c *UserActivityController) UpdateUserActivity(ctx *fiber.Ctx) error {
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

	// Call the UpdateUserActivity service method
	err := c.userActivityService.UpdateUserActivity(uint64(userId), param.StepId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to update user activity",
		}
	}

	return response.Ok(ctx, "user activity updated successfully")
}
