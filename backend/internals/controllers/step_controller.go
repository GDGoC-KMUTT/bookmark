package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type StepController struct {
	stepSvc services.StepService
}

func NewStepController(stepSvc services.StepService) StepController {
	return StepController{
		stepSvc: stepSvc,
	}
}

// GetStepInfo
// @ID getStepInfo
// @Tags step
// @Summary GetStepInfo
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Param q query payload.StepInfoQuery true "StepInfoQuery"
// @Success 200 {object} response.InfoResponse[[]payload.StepInfo]
// @Failure 400 {object} response.GenericError
// @Router /step/{stepId} [get]
func (r *StepController) GetStepInfo(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)
	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	query := new(payload.StepInfoQuery)
	if err := c.QueryParser(query); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepInfoQuery",
		}
	}

	// * validate body
	if err := utils.Validate.Struct(query); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		return &response.GenericError{
			Err: validationErrors,
		}
	}

	stepInfo, err := r.stepSvc.GetStepInfo(query.CourseId, query.ModuleId, param.StepId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get step info",
		}
	}

	return response.Ok(c, stepInfo)
}

// GetGemEachStep
// @ID getGemEachStep
// @Tags step
// @Summary GetGemEachStep
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[payload.GetGemsResponse]
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

	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	totalGems, currentGems, err := r.stepSvc.GetGems(param.StepId, &userId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to getGems",
		}
	}

	res := &payload.GetGemsResponse{
		TotalGems:   totalGems,
		CurrentGems: currentGems,
	}

	return response.Ok(c, res)
}

// GetStepComment
// @ID getStepComment
// @Tags step
// @Summary GetStepComment
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.StepCommentInfo]
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

	stepComments, err := r.stepSvc.GetStepComment(param.StepId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to getStepComment",
		}
	}

	return response.Ok(c, stepComments)
}

// CreateStepComment
// @ID createStepComment
// @Tags step
// @Summary CreateStepComment
// @Accept json
// @Produce json
// @Param q body payload.Comment true "Comment"
// @Success 200 {object} response.InfoResponse[string]
// @Failure 400 {object} response.GenericError
// @Router /step/comment/create [post]
func (r *StepController) CreateStepComment(c *fiber.Ctx) error {
	body := new(payload.Comment)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid comment payload",
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

	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	if err := r.stepSvc.CreteStpComment(body.StepId, &userId, body.Content); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to create stepComment",
		}
	}

	return response.Created(c, "successfully create step comment")
}

// CreateStepCommentUpVote
// @ID createStepCommentUpVote
// @Tags step
// @Summary CreateStepCommentUpVote
// @Accept json
// @Produce json
// @Param q body payload.UpVoteComment true "UpVoteComment"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/comment/upvote [post]
func (r *StepController) CreateStepCommentUpVote(c *fiber.Ctx) error {
	body := new(payload.UpVoteComment)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid comment payload",
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

	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	if err := r.stepSvc.CreateStepCommentUpVote(&userId, body.StepCommentId); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to create step comment upvote",
		}
	}

	return response.Created(c, "successfully create step comment upvote")
}

// GetStepEvaluate
// @ID getStepEvaluate
// @Tags step
// @Summary GetStepEvaluate
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.StepEvalInfo]
// @Failure 400 {object} response.GenericError
// @Router /step/stepEval/{stepId} [get]
func (r *StepController) GetStepEvaluate(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)

	// TODO
	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	stepEvals, err := r.stepSvc.GetStepEvalInfo(param.StepId)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get step eval info",
		}
	}

	return response.Ok(c, stepEvals)
}

// SubmitStepEval
// @ID submitStepEval
// @Tags step
// @Summary SubmitStepEval
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[[]payload.CourseWithFieldType]
// @Failure 400 {object} response.GenericError
// @Router /step/stepEval/{stepId} [post]
func (r *StepController) SubmitStepEval(c *fiber.Ctx) error {
	param := new(payload.StepIdParam)

	// TODO
	if err := c.ParamsParser(param); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid stepId param",
		}
	}

	res := new(payload.StepIdParam)
	return response.Ok(c, res)
}
