package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// ModuleController handles module-related endpoints
type ModuleController struct {
	moduleSvc services.ModuleServices
}

// NewModuleController initializes a new ModuleController
func NewModuleController(moduleSvc services.ModuleServices) *ModuleController {
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
	param := new(payload.ModuleIdParam)

	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid moduleId parameter",
		}
	}

	// Call service to get module info
	moduleInfo, err := c.moduleSvc.GetModuleInfo(strconv.FormatUint(*param.ModuleId, 10))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get module info",
		}
	}

	if moduleInfo == nil {
		return &response.GenericError{
			Err:     fmt.Errorf("module info is not available"),
			Message: "module info is not available",
		}
	}

	return response.Ok(ctx, moduleInfo)
}
