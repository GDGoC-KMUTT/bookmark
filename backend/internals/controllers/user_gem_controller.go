package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type GemController struct {
	gemSvc services.GemService
}

func NewGemController(gemSvc services.GemService) GemController {
	return GemController{
		gemSvc: gemSvc,
	}
}

// GetUserGems
// @ID getUserGems
// @Tags gems
// @Summary Fetch total gems for the user
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[payload.GemTotal]
// @Failure 400 {object} response.GenericError
// @Router /gems/total [get]
func (r *GemController) GetUserGems(c *fiber.Ctx) error {
	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	// * get total gems for user
	totalGems, err := r.gemSvc.GetTotalGems(uint(userId))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to fetch total gems",
		}
	}

	return response.Ok(c, totalGems)
}
