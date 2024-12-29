package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserStrengthController struct {
	userStrengthSvc services.UserStrengthService
}

func NewUserStrengthController(userStrengthSvc services.UserStrengthService) *UserStrengthController {
	return &UserStrengthController{
		userStrengthSvc: userStrengthSvc,
	}
}

// GetStrengthDataByUserID
// @ID getStrengthDataByUserID
// @Tags user-strength
// @Summary Get user strength data by user ID
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[response.StrengthDataResponse]
// @Failure 400 {object} response.GenericError
// @Router /strength/strength-info [get]
func (c *UserStrengthController) GetStrengthDataByUserID(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	strengthData, err := c.userStrengthSvc.GetStrengthDataByUserID(uint64(userId))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "Failed to fetch strength data",
		}
	}

	return response.Ok(ctx, strengthData)
}

// GetSuggestionCourse
// @ID getSuggestionCourse
// @Tags user-strength
// @Summary Get user strength data by user ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.InfoResponse[response.CourseSuggestionResponse]
// @Failure 400 {object} response.GenericError
// @Router /strength/suggestions [get]
func (c *UserStrengthController) GetSuggestionCourse(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	courses, err := c.userStrengthSvc.GetSuggestionCourse(uint64(userId))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "Failed to fetch course suggestions",
		}
	}

	return response.Ok(ctx, courses)
}
