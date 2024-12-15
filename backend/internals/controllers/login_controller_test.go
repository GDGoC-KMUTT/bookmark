package controllers

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/routes/handler"
	"backend/internals/services"
	"backend/internals/utils"
	mockServices "backend/mocks/services"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type LoginControllerTestSuite struct {
	suite.Suite
}

func setupTestLoginController(conf *config.Config, loginSvc services.LoginService) *fiber.App {
	config.BootConfiguration()

	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	loginController := NewLoginController(conf, loginSvc)

	app.Get("/login/redirect", loginController.LoginRedirect)
	app.Post("/login/callback", loginController.LoginCallBack)
	return app
}

func (suite *LoginControllerTestSuite) TestCallBackWhenSuccess() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	mockBodyReq := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockFirstName := utils.Ptr("fn")
	mockLastName := utils.Ptr("ln")
	mockEmail := utils.Ptr("test@gmail.com")
	mockProfileUrl := utils.Ptr("url")
	mockUserId := utils.Ptr[uint64](1)
	mockToken := utils.Ptr("signedToken")

	mockLoginService.EXPECT().OAuthSetup(mock.Anything).Return(&oidc.UserInfo{
		Subject:       *mockFirstName,
		Email:         *mockEmail,
		Profile:       *mockProfileUrl,
		EmailVerified: true,
	}, nil)

	mockLoginService.EXPECT().GetOrCreateUserFromClaims(mock.Anything).Return(&models.User{
		Id:        mockUserId,
		Email:     mockEmail,
		Firstname: mockFirstName,
		Lastname:  mockLastName,
		PhotoUrl:  mockProfileUrl,
	}, nil)

	mockLoginService.EXPECT().SignJwtToken(mock.Anything, mock.Anything).Return(mockToken, nil)

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/login/callback", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.CallbackResponse])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(*mockToken, *r.Data.Token)
	is.Equal(http.StatusOK, res.StatusCode)
}

func (suite *LoginControllerTestSuite) TestCallBackWhenFailedToParseBody() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	invalidBody := `{ "code": "somecode" `

	req := httptest.NewRequest(http.MethodPost, "/login/callback", strings.NewReader(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("Invalid JSON format in the request body", r.Message)
	is.Equal(http.StatusBadRequest, res.StatusCode)
}

func (suite *LoginControllerTestSuite) TestCallBackWhenValidateFailed() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	invalidBody := `{ "test": "somecode" }`

	req := httptest.NewRequest(http.MethodPost, "/login/callback", strings.NewReader(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("Request body validation failed", r.Message)
	is.Equal(http.StatusBadRequest, res.StatusCode)
}

func (suite *LoginControllerTestSuite) TestCallBackWhenFailedToSetupOAuth() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	mockBodyReq := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockLoginService.EXPECT().OAuthSetup(mock.Anything).Return(nil, fmt.Errorf("failed to setup oauth"))

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/login/callback", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("failed to setup OAuth", r.Message)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *LoginControllerTestSuite) TestCallBackWhenFailedToGetOrCreateFromClaim() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	mockBodyReq := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockFirstName := utils.Ptr("fn")
	mockEmail := utils.Ptr("test@gmail.com")
	mockProfileUrl := utils.Ptr("url")

	mockLoginService.EXPECT().OAuthSetup(mock.Anything).Return(&oidc.UserInfo{
		Subject:       *mockFirstName,
		Email:         *mockEmail,
		Profile:       *mockProfileUrl,
		EmailVerified: true,
	}, nil)

	mockLoginService.EXPECT().GetOrCreateUserFromClaims(mock.Anything).Return(nil, fmt.Errorf("failed to create from claim"))

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/login/callback", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("failed to get or create user from claims", r.Message)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *LoginControllerTestSuite) TestCallBackWhenFailedToSignJwt() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	mockBodyReq := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockFirstName := utils.Ptr("fn")
	mockLastName := utils.Ptr("ln")
	mockEmail := utils.Ptr("test@gmail.com")
	mockProfileUrl := utils.Ptr("url")
	mockUserId := utils.Ptr[uint64](1)

	mockLoginService.EXPECT().OAuthSetup(mock.Anything).Return(&oidc.UserInfo{
		Subject:       *mockFirstName,
		Email:         *mockEmail,
		Profile:       *mockProfileUrl,
		EmailVerified: true,
	}, nil)

	mockLoginService.EXPECT().GetOrCreateUserFromClaims(mock.Anything).Return(&models.User{
		Id:        mockUserId,
		Email:     mockEmail,
		Firstname: mockFirstName,
		Lastname:  mockLastName,
		PhotoUrl:  mockProfileUrl,
	}, nil)

	mockLoginService.EXPECT().SignJwtToken(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to sign jwt"))

	jsonBody, _ := json.Marshal(mockBodyReq)
	req := httptest.NewRequest(http.MethodPost, "/login/callback", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)

	r := new(response.GenericError)
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal("failed to sign jwt token", r.Message)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *LoginControllerTestSuite) TestLoginRedirectWhenSuccess() {
	is := assert.New(suite.T())

	mockLoginService := new(mockServices.LoginService)

	app := setupTestLoginController(config.Env, mockLoginService)

	req := httptest.NewRequest(http.MethodGet, "/login/redirect", nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.CallbackResponse])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(http.StatusFound, res.StatusCode)
}

func TestLoginController(t *testing.T) {
	suite.Run(t, new(LoginControllerTestSuite))
}
