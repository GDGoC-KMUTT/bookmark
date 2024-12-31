package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/entities/payload"
	"backend/internals/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
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
// @Success 200 {object} response.InfoResponse[[]payload.ModuleIdParam]
// @Failure 400 {object} response.GenericError
// @Failure 500 {object} response.GenericError
// @Router /step/{moduleId}/info [get]
func (c *ModuleStepController) GetModuleSteps(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["userId"].(float64))

	// Parse moduleId from path parameters
	param := new(payload.ModuleIdParam)
	if err := ctx.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid moduleId parameter",
		}
	}

	var moduleId string
    if param.ModuleId != nil {
        moduleId = strconv.FormatUint(*param.ModuleId, 10)
    } else {
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

	return response.Ok(ctx, steps)
}
