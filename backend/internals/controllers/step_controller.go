package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StepController struct {
	courseSvc services.CourseService
}

func NewStepController(courseSvc services.CourseService) StepController {
	return StepController{
		courseSvc: courseSvc,
	}
}

// GetStepInfo
// @ID getStepInfo
// @Tags step
// @Summary GetStepInfo
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/stepInfo/{stepId} [get]
func (r *StepController) GetStepInfo(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	res := new(payload.StepIdParam)
	return response.Ok(c, res)
}

// GetGemEachStep
// @ID getGemEachStep
// @Tags step
// @Summary GetGemEachStep
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/gem/{stepId} [get]
func (r *StepController) GetGemEachStep(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	res := new(payload.StepIdParam)
	return response.Ok(c, res)
}

// GetStepComment
// @ID getStepComment
// @Tags step
// @Summary GetStepComment
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/comment/{stepId} [get]
func (r *StepController) GetStepComment(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	res := new(payload.StepIdParam)
	return response.Ok(c, res)
}

// CreateStepComment
// @ID createStepComment
// @Tags step
// @Summary CreateStepComment
// @Accept json
// @Produce json
// @Param q body payload.OauthCallback true "OauthCallback"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/create [post]
func (r *StepController) CreateStepComment(c *fiber.Ctx) error {
	body := new(payload.StepIdParam)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	// * validate body
	if err := utils.Validate.Struct(body); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		return &response.GenericError{
			Err: validationErrors,
		}
	}

	res := new(payload.StepIdParam)
	return response.Ok(c, res)
}

// GetStepEvaluate
// @ID getStepEvaluate
// @Tags step
// @Summary GetStepEvaluate
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/stepEval/{stepId} [get]
func (r *StepController) GetStepEvaluate(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)

	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	res := new(payload.StepIdParam)
	return response.Ok(c, res)
}
