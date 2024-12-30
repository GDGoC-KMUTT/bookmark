package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ModuleStepController struct {
	moduleStepSvc services.ModuleStepServices
}

func NewModuleStepController(moduleStepSvc services.ModuleStepServices) *ModuleStepController {
	return &ModuleStepController{
		moduleStepSvc: moduleStepSvc,
	}
}

// GetModuleSteps
// @ID getModuleSteps
// @Tags moduleStep
// @Summary Get module steps with evaluation status
// @Description Fetch steps for a module and calculate evaluation status for each step.
// @Accept json
// @Produce json
// @Param moduleId path string true "Module ID"
// @Success 200 {object} response.InfoResponse[[]payload.ModuleStepResponse]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /step/{moduleId}/info [get]
func (c *ModuleStepController) GetModuleSteps(ctx *fiber.Ctx) error {
	// Extract user from JWT token using ctx.Locals
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["userId"].(float64))

	// Extract moduleId from URL parameter
	moduleId := ctx.Params("moduleId")
	if moduleId == "" {
		return &response.GenericError{
			Err:     fiber.NewError(fiber.StatusBadRequest, "module ID is required"),
			Message: "missing module ID",
		}
	}

	// Fetch steps from service
	steps, err := c.moduleStepSvc.GetModuleSteps(userId, moduleId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get module steps",
		}
	}

	// Return the steps as a response
	return response.Ok(ctx, steps)
}
