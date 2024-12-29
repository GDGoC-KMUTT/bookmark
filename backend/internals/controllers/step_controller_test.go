package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	"backend/internals/utils"
	mockServices "backend/mocks/services"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StepControllerTestSuit struct {
	suite.Suite
}

func setupTestStepController(mockStepService *mockServices.StepService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	// Initialize the controller
	stepController := NewStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

	mockBodyReq := &payload.Comment{
		StepId:  utils.Ptr(uint64(2)),
		Content: utils.Ptr("content"),
	}

	mockStepService.EXPECT().CreteStpComment(mock.Anything, mock.Anything, mock.Anything).Return(nil)

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

	app := setupTestStepController(mockStepService)

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

	app := setupTestStepController(mockStepService)

	mockBodyReq := &payload.Comment{
		StepId: utils.Ptr(uint64(2)),
	}

	mockStepService.EXPECT().CreteStpComment(mock.Anything, mock.Anything, mock.Anything).Return(nil)

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

	app := setupTestStepController(mockStepService)

	mockBodyReq := &payload.Comment{
		StepId:  utils.Ptr(uint64(2)),
		Content: utils.Ptr("content"),
	}

	mockStepService.EXPECT().CreteStpComment(mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("failed to createStepComment"))

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

	app := setupTestStepController(mockStepService)

	mockBodyReq := &payload.UpVoteComment{
		StepCommentId: utils.Ptr(uint64(3)),
	}

	mockStepService.EXPECT().CreateOrDeleteStepCommentUpVote(mock.Anything, mock.Anything).Return(nil)

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(string(jsonBody)))
	res, err := app.Test(req)

	r := new(response.InfoResponse[string])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal("failed", r.Message)
}

//func (suite *StepControllerTestSuit) TestUpVoteCommentWhenFailedToParseBody() {
//	is := assert.New(suite.T())
//
//	mockStepService := new(mockServices.StepService)
//
//	app := setupTestStepController(mockStepService)
//
//	mockBodyReq := "{body:"
//
//	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(mockBodyReq))
//	res, err := app.Test(req)
//
//	r := new(response.GenericError)
//	body, _ := io.ReadAll(res.Body)
//	json.Unmarshal(body, &r)
//
//	is.Nil(err)
//	is.Equal(http.StatusBadRequest, res.StatusCode)
//	is.Equal("Invalid JSON format in the request body", r.Message)
//}
//
//func (suite *StepControllerTestSuit) TestUpVoteCommentWhenValidationFailed() {
//	is := assert.New(suite.T())
//
//	mockStepService := new(mockServices.StepService)
//
//	app := setupTestStepController(mockStepService)
//
//	mockBodyReq := &payload.Comment{
//		StepId: utils.Ptr(uint64(2)),
//	}
//
//	jsonBody, _ := json.Marshal(mockBodyReq)
//	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(string(jsonBody)))
//	res, err := app.Test(req)
//
//	r := new(response.GenericError)
//	body, _ := io.ReadAll(res.Body)
//	json.Unmarshal(body, &r)
//
//	is.Nil(err)
//	is.Equal(http.StatusBadRequest, res.StatusCode)
//	is.Equal("Request body validation failed", r.Message)
//}
//
//func (suite *StepControllerTestSuit) TestUpVoteCommentWhenFailedToCreateOrDeleteStepCommentUpVote() {
//	is := assert.New(suite.T())
//
//	mockStepService := new(mockServices.StepService)
//
//	app := setupTestStepController(mockStepService)
//
//	mockBodyReq := &payload.Comment{
//		StepId:  utils.Ptr(uint64(2)),
//		Content: utils.Ptr("content"),
//	}
//
//	mockStepService.EXPECT().CreateOrDeleteStepCommentUpVote(mock.Anything, mock.Anything).Return(fmt.Errorf("failed to upvote"))
//
//	jsonBody, _ := json.Marshal(mockBodyReq)
//	req := httptest.NewRequest(http.MethodPost, "/step/comment/upvote", strings.NewReader(string(jsonBody)))
//	res, err := app.Test(req)
//
//	r := new(response.GenericError)
//	body, _ := io.ReadAll(res.Body)
//	json.Unmarshal(body, &r)
//
//	is.Nil(err)
//	is.Equal(http.StatusInternalServerError, res.StatusCode)
//	is.Equal("failed to create step comment upvote", r.Message)
//}

func TestStepController(t *testing.T) {
	suite.Run(t, new(StepControllerTestSuit))
}
