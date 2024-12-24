package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ProfileController struct {
	profileSvc services.ProfileService
}

func NewProfileController(profileSvc services.ProfileService) ProfileController {
	return ProfileController{
		profileSvc: profileSvc,
	}
}

// ProfileUserInfo
// @ID profileUserInfo
// @Tags profile
// @Summary profileUserInfo
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[payload.Profile]
// @Failure 400 {object} response.GenericError
// @Router /profile/info [get]
func (r *ProfileController) ProfileUserInfo(c *fiber.Ctx) error {
	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	// * get user profile
	userProfile, err := r.profileSvc.GetUserInfo(utils.Ptr(strconv.Itoa(int(userId))))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get user profile",
		}
	}

	return response.Ok(c, userProfile)
}

// GetUserGems
// @ID getUserGems
// @Tags gems
// @Summary Fetch total gems for the user
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[payload.GemTotal]
// @Failure 400 {object} response.GenericError
// @Router /profile/totalgems [get]
func (r *ProfileController) GetUserGems(c *fiber.Ctx) error {
	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, ok := claims["userId"].(float64)
	if !ok {
		return &response.GenericError{
			Err:     fmt.Errorf("invalid user ID"),
			Message: "Invalid user ID in JWT token",
		}
	}
	fmt.Println("userId:", userId)
	
	// * get total gems for user
	totalGems, err := r.profileSvc.GetTotalGems(uint(userId))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to fetch total gems",
		}
	}

	return response.Ok(c, totalGems)
}