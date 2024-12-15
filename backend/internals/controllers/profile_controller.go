package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
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
