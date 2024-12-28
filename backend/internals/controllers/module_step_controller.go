package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
)

// ModuleStepController handles module step-related endpoints
type ModuleStepController struct {
	moduleStepSvc *services.ModuleStepService
}

// NewModuleStepController initializes a new ModuleStepController
func NewModuleStepController(moduleStepSvc *services.ModuleStepService) *ModuleStepController {
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
	moduleId := ctx.Params("moduleId")

	// Call service to get module steps
	steps, err := c.moduleStepSvc.GetModuleSteps(moduleId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get module steps",
		})
	}

	// Return the response
	return ctx.JSON(&response.InfoResponse[[]payload.ModuleStepResponse]{
		Data: steps,
	})
}
