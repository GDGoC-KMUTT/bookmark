package services_test

import (
	"backend/internals/db/models"
	"backend/internals/entities/payload"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	mockUtilServices "backend/mocks/utils"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"testing"
	"time"
)

type LoginServiceTestSuite struct {
	suite.Suite
}

func (suite *LoginServiceTestSuite) TestOAuthSetupWhenSuccess() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuthService := new(mockUtilServices.OAuthService)
	mockJwtService := new(mockUtilServices.Jwt)

	mockBody := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockToken := &oauth2.Token{
		AccessToken:  "accessToken",
		Expiry:       time.Now().Add(24 * time.Hour),
		RefreshToken: "refreshToken",
		TokenType:    "type",
	}

	mockUserInfo := &oidc.UserInfo{
		Email:         "test@gmail.com",
		Profile:       "url",
		EmailVerified: true,
		Subject:       "test",
	}

	mockOAuthService.EXPECT().Exchange(mock.Anything, mock.Anything).Return(mockToken, nil)
	mockOAuthService.EXPECT().UserInfo(mock.Anything, mock.Anything).Return(mockUserInfo, nil)

	underTest := services.NewLoginService(mockUserRepo, mockOAuthService, mockJwtService)

	// Test Success
	result, err := underTest.OAuthSetup(mockBody)

	is.Nil(err)
	is.NotNil(result)
}

func (suite *LoginServiceTestSuite) TestOAuthSetupWhenExchangeFailed() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuthService := new(mockUtilServices.OAuthService)
	mockJwtService := new(mockUtilServices.Jwt)

	mockBody := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockOAuthService.EXPECT().Exchange(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to exchange"))

	underTest := services.NewLoginService(mockUserRepo, mockOAuthService, mockJwtService)

	// Test Success
	result, err := underTest.OAuthSetup(mockBody)

	is.Nil(result)
	is.NotNil(err)
	is.Equal("failed to exchange", err.Error())
}

func (suite *LoginServiceTestSuite) TestOAuthSetupWhenFailedToGetUserInfo() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuthService := new(mockUtilServices.OAuthService)
	mockJwtService := new(mockUtilServices.Jwt)

	mockBody := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockToken := &oauth2.Token{
		AccessToken:  "accessToken",
		Expiry:       time.Now().Add(24 * time.Hour),
		RefreshToken: "refreshToken",
		TokenType:    "type",
	}

	mockOAuthService.EXPECT().Exchange(mock.Anything, mock.Anything).Return(mockToken, nil)
	mockOAuthService.EXPECT().UserInfo(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get user info"))

	underTest := services.NewLoginService(mockUserRepo, mockOAuthService, mockJwtService)

	// Test Success
	result, err := underTest.OAuthSetup(mockBody)

	is.Nil(result)
	is.NotNil(err)
	is.Equal("failed to get user info", err.Error())
}

func (suite *LoginServiceTestSuite) TestSignJwtTokenWhenSuccess() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuthService := new(mockUtilServices.OAuthService)
	mockJwtService := new(mockUtilServices.Jwt)

	mockUser := &models.User{
		Id:        utils.Ptr[uint64](1),
		Email:     utils.Ptr("test@gmail.com"),
		Lastname:  utils.Ptr("ln"),
		PhotoUrl:  utils.Ptr("url"),
		Firstname: utils.Ptr("fn"),
	}

	mockSecret := utils.Ptr("super-secret")

	mockJwtService.EXPECT().NewWithClaims(mock.Anything, mock.Anything).Return(&jwt.Token{})
	mockJwtService.EXPECT().SignedString(mock.Anything, mock.Anything).Return("signedToken", nil)

	underTest := services.NewLoginService(mockUserRepo, mockOAuthService, mockJwtService)

	// Test Success
	result, err := underTest.SignJwtToken(mockUser, mockSecret)

	is.Nil(err)
	is.NotNil(result)
	is.Equal("signedToken", *result)
}

func (suite *LoginServiceTestSuite) TestSignJwtTokenWhenSignStringError() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuthService := new(mockUtilServices.OAuthService)
	mockJwtService := new(mockUtilServices.Jwt)

	mockUser := &models.User{
		Id:        utils.Ptr[uint64](1),
		Email:     utils.Ptr(""),
		Lastname:  utils.Ptr(""),
		PhotoUrl:  utils.Ptr(""),
		Firstname: utils.Ptr(""),
	}

	mockSecret := utils.Ptr("super-secret")

	mockJwtService.EXPECT().NewWithClaims(mock.Anything, mock.Anything).Return(&jwt.Token{})
	mockJwtService.EXPECT().SignedString(mock.Anything, mock.Anything).Return("", fmt.Errorf("failed to signed string"))

	underTest := services.NewLoginService(mockUserRepo, mockOAuthService, mockJwtService)

	// Test Success
	_, err := underTest.SignJwtToken(mockUser, mockSecret)

	is.NotNil(err)
	is.Equal("failed to signed string", err.Error())
}

func (suite *LoginServiceTestSuite) TestGetOrCreateUserFromClaimWhenClaimNotSetError() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuthService := new(mockUtilServices.OAuthService)
	mockJwtService := new(mockUtilServices.Jwt)

	mockUserInfo := &oidc.UserInfo{
		Email:         "test@gmail.com",
		Profile:       "profile",
		Subject:       "sub",
		EmailVerified: true,
	}

	underTest := services.NewLoginService(mockUserRepo, mockOAuthService, mockJwtService)
	// Test Success
	result, err := underTest.GetOrCreateUserFromClaims(mockUserInfo)

	is.NotNil(err)
	is.Nil(result)
	is.Equal("oidc: claims not set", err.Error())
}

func TestLoginService(t *testing.T) {
	suite.Run(t, new(LoginServiceTestSuite))
}
