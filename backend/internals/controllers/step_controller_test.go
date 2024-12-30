package controllers

import (
	"backend/internals/config"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	"backend/internals/utils"
	mockServices "backend/mocks/services"
	mockUtilServices "backend/mocks/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StepControllerTestSuit struct {
	suite.Suite
}

func setupTestStepController(mockStepService *mockServices.StepService, mockMinioService *mockUtilServices.MinioService) *fiber.App {
	config.BootConfiguration()

	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	// Initialize the controller
	stepController := NewStepController(mockStepService, config.Env, mockMinioService)

	// Middleware to simulate JWT Locals
	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{"userId": float64(123)} // Simulate a valid userId claim
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	// Register the route
	step := app.Group("/step")
	step.Get("/:stepId", stepController.GetStepInfo)
	step.Get("/gem/:stepId", stepController.GetGemEachStep)

	stepEval := step.Group("/stepEval")
	stepEval.Post("/submit", stepController.SubmitStepEval)
	stepEval.Get("/status", stepController.CheckStepEvalStatus)
	stepEval.Post("/submit-type-check", stepController.SubmitStepEvalTypCheck)
	stepEval.Get("/:stepId", stepController.GetStepEvaluate)

	stepComment := step.Group("/comment")
	stepComment.Post("/create", stepController.CommentOnStep)
	stepComment.Post("/upvote", stepController.UpVoteStepComment)
	stepComment.Get("/:stepId", stepController.GetStepComment)
	return app
}
func (suite *StepControllerTestSuit) TestGetStepInfoWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	mockStepDetail := &payload.StepDetail{
		StepId:      mockStepId,
		Content:     utils.Ptr("content"),
		Banner:      utils.Ptr("banner"),
		Check:       utils.Ptr("check"),
		Error:       utils.Ptr("error"),
		Description: utils.Ptr("desc"),
		ModuleId:    utils.Ptr(uint64(2)),
		Outcome:     utils.Ptr("outcome"),
		Title:       utils.Ptr("title"),
	}

	mockAuthors := []*payload.UserInfo{
		{
			UserId:    utils.Ptr(uint64(1)),
			Email:     utils.Ptr("email"),
			FirstName: utils.Ptr("fn"),
			LastName:  utils.Ptr("ln"),
			PhotoUrl:  utils.Ptr("photo"),
		},
	}

	mockUserPassed := []*payload.UserInfo{
		{
			UserId:    utils.Ptr(uint64(2)),
			Email:     utils.Ptr("email"),
			FirstName: utils.Ptr("fn"),
			LastName:  utils.Ptr("ln"),
			PhotoUrl:  utils.Ptr("photo"),
		},
	}

	mockStepService.EXPECT().GetStepInfo(mock.Anything).Return(&payload.StepInfo{
		Step:       mockStepDetail,
		Authors:    mockAuthors,
		UserPassed: mockUserPassed,
	}, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.StepInfo])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestGetStepInfoWhenFailedToParseParam() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	req := httptest.NewRequest(http.MethodGet, "/step/t", nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("invalid stepId param", r.Message)
}

func (suite *StepControllerTestSuit) TestGetStepInfoWhenFailedToGetStepInfo() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	mockStepService.EXPECT().GetStepInfo(mock.Anything).Return(nil, fmt.Errorf("failed to get stepInfo"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("failed to get step info", r.Message)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestGetGemWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))
	mockTotalGem := utils.Ptr(5)
	mockCurrentGem := utils.Ptr(2)

	mockStepService.EXPECT().GetGems(mock.Anything, mock.Anything).Return(mockTotalGem, mockCurrentGem, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/gem/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.GetGemsResponse])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestGetGemWhenFailedToParseParam() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/gem/:%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("invalid stepId param", r.Message)
}

func (suite *StepControllerTestSuit) TestGetGemWhenFailedToGetGem() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	mockStepService.EXPECT().GetGems(mock.Anything, mock.Anything).Return(nil, nil, fmt.Errorf("failed to get gem"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/gem/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to getGems", r.Message)
}

func (suite *StepControllerTestSuit) TestGetStepCommentWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	mockStepComment := []payload.StepCommentInfo{
		{
			Comment:       utils.Ptr("comment"),
			StepCommentId: utils.Ptr(uint64(2)),
		},
	}

	mockStepService.EXPECT().GetStepComment(mock.Anything, mock.Anything).Return(mockStepComment, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/comment/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[[]payload.StepCommentInfo])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestGetStepCommentWhenFailedToParseParam() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/comment/:%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("invalid stepId param", r.Message)
}

func (suite *StepControllerTestSuit) TestGetStepCommentWhenFailedToGetStepComment() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(2))

	mockStepService.EXPECT().GetStepComment(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get step comment"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/comment/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to getStepComment", r.Message)
}

func (suite *StepControllerTestSuit) TestCommentOnStepWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.Comment{
		StepId:  utils.Ptr(uint64(2)),
		Content: utils.Ptr("content"),
	}

	mockStepService.EXPECT().CreateStpComment(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/create", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.InfoResponse[string])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestCommentOnStepWhenFailedToParseBody() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := "{body:"

	req := httptest.NewRequest(http.MethodPost, "/step/comment/create", strings.NewReader(mockBodyReq))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Invalid JSON format in the request body", r.Message)
}

func (suite *StepControllerTestSuit) TestCommentOnStepWhenValidationFailed() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.Comment{
		StepId: utils.Ptr(uint64(2)),
	}

	mockStepService.EXPECT().CreateStpComment(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/create", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Request body validation failed", r.Message)
}

func (suite *StepControllerTestSuit) TestCommentOnStepWhenFailedToCreateStepComment() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.Comment{
		StepId:  utils.Ptr(uint64(2)),
		Content: utils.Ptr("content"),
	}

	mockStepService.EXPECT().CreateStpComment(mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("failed to createStepComment"))

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/create", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to create stepComment", r.Message)
}

func (suite *StepControllerTestSuit) TestUpVoteCommentWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.UpVoteComment{
		StepCommentId: utils.Ptr(uint64(3)),
	}

	mockStepService.EXPECT().CreateOrDeleteStepCommentUpVote(mock.Anything, mock.Anything).Return(nil)

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	r := new(response.InfoResponse[string])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestUpVoteCommentWhenFailedToParseBody() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := "{body:"

	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(mockBodyReq))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Invalid JSON format in the request body", r.Message)
}

func (suite *StepControllerTestSuit) TestUpVoteCommentWhenValidationFailed() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.UpVoteComment{}

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Request body validation failed", r.Message)
}

func (suite *StepControllerTestSuit) TestUpVoteCommentWhenFailedToCreateOrDeleteStepCommentUpVote() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.UpVoteComment{
		StepCommentId: utils.Ptr(uint64(3)),
	}

	mockStepService.EXPECT().CreateOrDeleteStepCommentUpVote(mock.Anything, mock.Anything).Return(fmt.Errorf("failed to upvote"))

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to create step comment upvote", r.Message)
}

func (suite *StepControllerTestSuit) TestGetStepEvalWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(3))

	mockStepEvals := []*payload.StepEvalInfo{
		{
			StepId: mockStepId,
		},
	}

	mockStepService.EXPECT().GetStepEvalInfo(mock.Anything, mock.Anything).Return(mockStepEvals, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/stepEval/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[[]payload.GetGemsResponse])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestGetStepEvalWhenFailedToParseParam() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(3))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/stepEval/:%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("invalid stepId param", r.Message)
}

func (suite *StepControllerTestSuit) TestGetStepEvalWhenFailedToGetStepEval() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepId := utils.Ptr(uint64(3))

	mockStepService.EXPECT().GetStepEvalInfo(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get stepEval"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/stepEval/%d", mockStepId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeCheckWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.StepEvalIdBody{
		StepEvalId: utils.Ptr(uint64(3)),
	}

	mockUserEvalId := utils.Ptr(uint64(2))

	mockStepService.EXPECT().SubmitStepEvalTypeCheck(mock.Anything, mock.Anything).Return(mockUserEvalId, nil)

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/stepEval/submit-type-check", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.CreateUserEvalRes])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeCheckWhenFailedToParseBody() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := "{body:"

	req := httptest.NewRequest(http.MethodPost, "/step/stepEval/submit-type-check", strings.NewReader(mockBodyReq))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Invalid JSON format in the request body", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeCheckWhenValidationFailed() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.StepEvalIdBody{}

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/stepEval/submit-type-check", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Request body validation failed", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeCheckWhenFailedToSubmit() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockBodyReq := &payload.StepEvalIdBody{
		StepEvalId: utils.Ptr(uint64(3)),
	}

	mockStepService.EXPECT().SubmitStepEvalTypeCheck(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to submit "))

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/stepEval/submit-type-check", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to submit step eval type check", r.Message)
}

func (suite *StepControllerTestSuit) TestCheckStepEvalStatusWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockUserEvalId := utils.Ptr(uint64(2))

	mockUserEvalResult := &payload.UserEvalResult{
		Content: utils.Ptr("content"),
		Comment: utils.Ptr("comment"),
	}

	mockStepService.EXPECT().CheckStepEvalStatus(mock.Anything, mock.Anything).Return(mockUserEvalResult, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/stepEval/status?userEvalId=%d", mockUserEvalId), nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.UserEvalResult])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestCheckStepEvalStatusWheFailedToParseQuery() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	req := httptest.NewRequest(http.MethodGet, "/step/stepEval/status?userEvalId=invalid", nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("invalid userEvalId query request", r.Message)
}

func (suite *StepControllerTestSuit) TestCheckStepEvalStatusWhenValidationFailed() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	req := httptest.NewRequest(http.MethodGet, "/step/stepEval/status?", nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("Request body validation failed", r.Message)
}

func (suite *StepControllerTestSuit) TestCheckStepEvalStatusWhenFailedToGetStatus() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockUserEvalId := utils.Ptr(uint64(2))

	mockStepService.EXPECT().CheckStepEvalStatus(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get status"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/step/stepEval/status?userEvalId=%d", mockUserEvalId), nil)
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get latest status of each user eval", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeTextWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockUserEvalId := utils.Ptr(uint64(1))

	mockStepService.EXPECT().CreateUserEval(mock.Anything).Return(mockUserEvalId, nil)

	formData := "data={\"stepId\":1, \"stepEvalId\":123, \"content\": \"Valid content\"}"
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.CreateUserEvalRes])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeTextWhenFailedToParseJson() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockUserEvalId := utils.Ptr(uint64(1))

	mockStepService.EXPECT().CreateUserEval(mock.Anything).Return(mockUserEvalId, nil)

	formData := "data={\"stepId\":1, \"stepEvalId\":123, \"content\": \"Valid content\""
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Invalid JSON format in the request body", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeTextWhenValidationFailed() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockUserEvalId := utils.Ptr(uint64(1))

	mockStepService.EXPECT().CreateUserEval(mock.Anything).Return(mockUserEvalId, nil)

	formData := "data={\"stepId\":1, \"content\": \"Valid content\"}"
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Request body validation failed", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalWhenFailedToCreateUserEval() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepService.EXPECT().CreateUserEval(mock.Anything).Return(nil, fmt.Errorf("failed to creat userEval"))

	formData := "data={\"stepId\":1, \"stepEvalId\":123, \"content\": \"Valid content\"}"
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to create user eval", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeImageWhenSuccess() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockUserEvalId := utils.Ptr(uint64(1))
	mockFileName := utils.Ptr("file.png")

	mockStepService.EXPECT().CreateFileFormat(mock.Anything, mock.Anything, mock.Anything).Return(mockFileName, nil)
	mockStepService.EXPECT().CreateUserEval(mock.Anything).Return(mockUserEvalId, nil)
	mockMinioService.EXPECT().PutObject(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Prepare the form with the JSON data and file
	formData := new(bytes.Buffer)
	writer := multipart.NewWriter(formData)

	// Add the JSON data (e.g., for "data" field)
	jsonPart, _ := writer.CreateFormField("data")
	jsonPart.Write([]byte("{\"stepId\":1, \"stepEvalId\":123}"))

	// Create a dummy file and add it to the form
	filePart, _ := writer.CreateFormFile("file", "test.png")
	filePart.Write([]byte("This is a dummy file content"))

	// Close the writer to finalize the form
	writer.Close()

	// Create a new request with the form data
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", formData)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Sending the request and getting the response
	res, err := app.Test(req)

	// Read the response body and unmarshal it
	r := new(response.InfoResponse[payload.CreateUserEvalRes])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	// Assert the expected results
	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeImageWhenFileNotFound() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	// Prepare the form with the JSON data and file
	formData := new(bytes.Buffer)
	writer := multipart.NewWriter(formData)

	// Add the JSON data (e.g., for "data" field)
	jsonPart, _ := writer.CreateFormField("data")
	jsonPart.Write([]byte("{\"stepId\":1, \"stepEvalId\":123}"))

	// Create a dummy file and add it to the form
	filePart, _ := writer.CreateFormFile("file", "")
	filePart.Write([]byte("This is a dummy file content"))

	// Close the writer to finalize the form
	writer.Close()

	// Create a new request with the form data
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", formData)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Sending the request and getting the response
	res, err := app.Test(req)

	// Read the response body and unmarshal it
	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	// Assert the expected results
	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("file not found", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeImageWhenFailedToCreateFileFormat() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockStepService.EXPECT().CreateFileFormat(mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to createFileFormat"))

	// Prepare the form with the JSON data and file
	formData := new(bytes.Buffer)
	writer := multipart.NewWriter(formData)

	// Add the JSON data (e.g., for "data" field)
	jsonPart, _ := writer.CreateFormField("data")
	jsonPart.Write([]byte("{\"stepId\":1, \"stepEvalId\":123}"))

	// Create a dummy file and add it to the form
	filePart, _ := writer.CreateFormFile("file", "test.png")
	filePart.Write([]byte("This is a dummy file content"))

	// Close the writer to finalize the form
	writer.Close()

	// Create a new request with the form data
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", formData)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Sending the request and getting the response
	res, err := app.Test(req)

	// Read the response body and unmarshal it
	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	// Assert the expected results
	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to create file format", r.Message)
}

func (suite *StepControllerTestSuit) TestSubmitStepEvalTypeImageWhenFailedToPutObject() {
	is := assert.New(suite.T())

	mockStepService := new(mockServices.StepService)
	mockMinioService := new(mockUtilServices.MinioService)

	app := setupTestStepController(mockStepService, mockMinioService)

	mockFileName := utils.Ptr("file.png")

	mockStepService.EXPECT().CreateFileFormat(mock.Anything, mock.Anything, mock.Anything).Return(mockFileName, nil)
	mockMinioService.EXPECT().PutObject(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("failed to put object"))

	// Prepare the form with the JSON data and file
	formData := new(bytes.Buffer)
	writer := multipart.NewWriter(formData)

	// Add the JSON data (e.g., for "data" field)
	jsonPart, _ := writer.CreateFormField("data")
	jsonPart.Write([]byte("{\"stepId\":1, \"stepEvalId\":123}"))

	// Create a dummy file and add it to the form
	filePart, _ := writer.CreateFormFile("file", "test.png")
	filePart.Write([]byte("This is a dummy file content"))

	// Close the writer to finalize the form
	writer.Close()

	// Create a new request with the form data
	req := httptest.NewRequest(fiber.MethodPost, "/step/stepEval/submit", formData)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Sending the request and getting the response
	res, err := app.Test(req)

	// Read the response body and unmarshal it
	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	// Assert the expected results
	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to upload file", r.Message)
}

func TestStepController(t *testing.T) {
	suite.Run(t, new(StepControllerTestSuit))
}
