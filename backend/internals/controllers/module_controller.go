package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
)

// ModuleController handles module-related endpoints
type ModuleController struct {
	moduleSvc *services.ModuleService
}

// NewModuleController initializes a new ModuleController
func NewModuleController(moduleSvc *services.ModuleService) *ModuleController {
	return &ModuleController{
		moduleSvc: moduleSvc,
	}
}

// GetModuleInfo
// @ID getModuleInfo
// @Tags module
// @Summary Get module information
// @Accept json
// @Produce json
// @Param moduleId path string true "Module ID"
// @Success 200 {object} response.InfoResponse[payload.ModuleResponse]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /module/{moduleId}/info [get]
func (c *ModuleController) GetModuleInfo(ctx *fiber.Ctx) error {
	moduleId := ctx.Params("moduleId")

	// Call service to get module info
	moduleInfo, err := c.moduleSvc.GetModuleInfo(moduleId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&response.GenericError{
			Err:     err,
			Message: "failed to get module info",
		})
	}

	// Dereference moduleInfo to pass the value to InfoResponse
	return ctx.JSON(&response.InfoResponse[payload.ModuleResponse]{
		Data: *moduleInfo, // Dereference the pointer here
	})
}
