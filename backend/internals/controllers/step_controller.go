package controllers

import (
	"backend/internals/config"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	"backend/internals/utils"
	utilServices "backend/internals/utils/services"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/url"
)

type StepController struct {
	stepSvc      services.StepService
	conf         *config.Config
	minioService utilServices.MinioService
}

func NewStepController(stepSvc services.StepService, conf *config.Config, minioService utilServices.MinioService) StepController {
	return StepController{
		stepSvc:      stepSvc,
		conf:         conf,
		minioService: minioService,
	}
}

// GetStepInfo
// @ID getStepInfo
// @Tags step
// @Summary GetStepInfo
// @Accept json
// @Produce json
// @Param stepId path uint true "Step ID"
// @Success 200 {object} response.InfoResponse[payload.StepInfo]
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

	stepInfo, err := r.stepSvc.GetStepInfo(param.StepId)
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

	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	stepComments, err := r.stepSvc.GetStepComment(param.StepId, utils.Ptr(uint64(userId)))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to getStepComment",
		}
	}

	return response.Ok(c, stepComments)
}

// CommentOnStep
// @ID commentOnStep
// @Tags step
// @Summary CommentOnStep
// @Accept json
// @Produce json
// @Param q body payload.Comment true "Comment"
// @Success 200 {object} response.InfoResponse[string]
// @Failure 400 {object} response.GenericError
// @Router /step/comment/create [post]
func (r *StepController) CommentOnStep(c *fiber.Ctx) error {
	body := new(payload.Comment)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err: err,
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

	if err := r.stepSvc.CreateStpComment(body.StepId, &userId, body.Content); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to create stepComment",
		}
	}

	return response.Created(c, "successfully create step comment")
}

// UpVoteStepComment
// @ID upVoteStepComment
// @Tags step
// @Summary UpVoteStepComment
// @Accept json
// @Produce json
// @Param q body payload.UpVoteComment true "UpVoteComment"
// @Success 200 {object} response.InfoResponse[string]
// @Failure 400 {object} response.GenericError
// @Router /step/comment/upvote [post]
func (r *StepController) UpVoteStepComment(c *fiber.Ctx) error {
	body := new(payload.UpVoteComment)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err: err,
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

	if err := r.stepSvc.CreateOrDeleteStepCommentUpVote(&userId, body.StepCommentId); err != nil {
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

	stepEvals, err := r.stepSvc.GetStepEvalInfo(param.StepId, &userId)
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
// @Param data formData string true "JSON data as string"
// @Param file formData file false "File to upload"
// @Success 200 {object} response.InfoResponse[payload.CreateUserEvalRes]
// @Failure 400 {object} response.GenericError
// @Router /step/stepEval/submit [post]
func (r *StepController) SubmitStepEval(c *fiber.Ctx) error {
	// Parse JSON from "data" form field
	jsonData := c.FormValue("data")
	body := new(payload.SubmitStepEval)
	if err := json.Unmarshal([]byte(jsonData), body); err != nil {
		return &response.GenericError{
			Err: err,
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

	userEval := &payload.CreateUserEvalReq{
		UserId:     &userId,
		StepEvalId: body.StepEvalId,
		Content:    body.Content,
	}

	result := &payload.CreateUserEvalRes{
		Pass: utils.Ptr(false),
	}

	result.UserSubmission = body.Content

	if body.Content == nil {
		// * Parse file form
		// Note: file is a *multipart.FileHeader instance
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return &response.GenericError{
				Err:     err,
				Message: "file not found",
			}
		}

		// * Convert multipart.FileHeader to File (which satisfies io.Reader)
		// Note: Since file is a *multipart.FileHeader instance
		// and minio.PutObject() requires a io.Reader instance
		file, err := fileHeader.Open()
		if err != nil {
			return &response.GenericError{
				Err:     err,
				Message: "failed to open file",
			}
		}

		// * Generate filename
		filename, err := r.stepSvc.CreateFileFormat(body.StepId, body.StepEvalId, &userId)
		if err != nil {
			return &response.GenericError{
				Err:     err,
				Message: "failed to create file format",
			}
		}

		userEval.Content = filename

		content, err := url.JoinPath(*config.Env.MinioS3Endpoint, *config.Env.MinioS3BucketName, *filename)
		if err != nil {
			return err
		}
		result.UserSubmission = &content

		// * Upload file to minio
		if err := r.minioService.PutObject(
			c.Context(),
			*r.conf.MinioS3BucketName,
			*filename,
			file,
			fileHeader,
		); err != nil {
			return &response.GenericError{
				Err:     err,
				Message: "failed to upload file",
			}
		}

	}

	userStepEvalId, err := r.stepSvc.CreateUserEval(userEval)
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to create user eval",
		}
	}

	result.UserEvalId = userStepEvalId

	return response.Ok(c, result)

}

// CheckStepEvalStatus
// @ID checkStepEvalStatus
// @Tags step
// @Summary CheckStepEvalStatus
// @Accept json
// @Produce json
// @Param q query payload.UserEvalIdsBody true "UserEvalIdsBody"
// @Success 200 {object} response.InfoResponse[payload.UserEvalResult]
// @Failure 400 {object} response.GenericError
// @Router /step/stepEval/status [get]
func (r *StepController) CheckStepEvalStatus(c *fiber.Ctx) error {
	query := new(payload.UserEvalIdsBody)

	if err := c.QueryParser(query); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "invalid userEvalId query request",
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

	// * login state
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	userEval, err := r.stepSvc.CheckStepEvalStatus(query.UserEvalId, utils.Ptr(uint64(userId)))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to get latest status of each user eval",
		}
	}

	return response.Ok(c, userEval)
}

// SubmitStepEvalTypCheck
// @ID submitStepEvalTypCheck
// @Tags step
// @Summary SubmitStepEvalTypCheck
// @Accept json
// @Produce json
// @Param q body payload.StepEvalIdBody true "StepEvalIdBody"
// @Success 200 {object} response.InfoResponse[payload.CreateUserEvalRes]
// @Failure 400 {object} response.GenericError
// @Router /step/stepEval/submit-type-check [post]
func (r *StepController) SubmitStepEvalTypCheck(c *fiber.Ctx) error {
	body := new(payload.StepEvalIdBody)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to parse body",
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

	userEvalId, err := r.stepSvc.SubmitStepEvalTypeCheck(body.StepEvalId, utils.Ptr(uint64(userId)))
	if err != nil {
		return &response.GenericError{
			Err:     err,
			Message: "failed to submit step eval type check",
		}
	}

	result := &payload.CreateUserEvalRes{
		UserEvalId: userEvalId,
	}

	return response.Ok(c, result)
}
