package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// ModuleStepController handles module step-related endpoints
type ModuleStepController struct {
	moduleStepSvc services.ModuleStepServices
}

// NewModuleStepController initializes a new ModuleStepController
func NewModuleStepController(moduleStepSvc services.ModuleStepServices) *ModuleStepController {
	return &ModuleStepController{
		moduleStepSvc: moduleStepSvc,
	}
}

// GetModuleSteps
// @ID getModuleSteps
// @Tags moduleStep
// @Summary Get all steps for a module
// @Accept json
// @Produce json
// @Param moduleId path string true "Module ID"
// @Success 200 {object} response.InfoResponse[[]payload.ModuleStepResponse]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /step/{moduleId}/info [get]
func (c *ModuleStepController) GetModuleSteps(ctx *fiber.Ctx) error {
	param := new(payload.ModuleIdParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid moduleId parameter",
		}
	}

	// Call service to get module steps
	steps, err := c.moduleStepSvc.GetModuleSteps(strconv.FormatUint(*param.ModuleId, 10))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get module steps",
		}
	}

	return response.Ok(ctx, steps)
}
