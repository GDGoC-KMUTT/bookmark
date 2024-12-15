package services_test

import (
	"backend/internals/entities/payload"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	mockServices "backend/mocks/services"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
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
	mockOAuth2Client := new(mockServices.OAuth2Client)
	mockOIdcProvider := new(mockServices.OIDCProvider)

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

	mockOAuth2Client.EXPECT().Exchange(mock.Anything, mock.Anything).Return(mockToken, nil)
	mockOIdcProvider.EXPECT().UserInfo(mock.Anything, mock.Anything).Return(mockUserInfo, nil)

	underTest := services.NewLoginService(mockUserRepo, mockOAuth2Client, mockOIdcProvider)

	// Test Success
	result, err := underTest.OAuthSetup(mockBody)

	is.Nil(err)
	is.NotNil(result)
}

func (suite *LoginServiceTestSuite) TestOAuthSetupWhenExchangeFailed() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuth2Client := new(mockServices.OAuth2Client)
	mockOIdcProvider := new(mockServices.OIDCProvider)

	mockBody := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockOAuth2Client.EXPECT().Exchange(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to exchange"))

	underTest := services.NewLoginService(mockUserRepo, mockOAuth2Client, mockOIdcProvider)

	// Test Success
	result, err := underTest.OAuthSetup(mockBody)

	is.Nil(result)
	is.NotNil(err)
	is.ErrorIs(err, fmt.Errorf("failed to exchange"))
}

func (suite *LoginServiceTestSuite) TestOAuthSetupWhenFailedToGetUserInfo() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)
	mockOAuth2Client := new(mockServices.OAuth2Client)
	mockOIdcProvider := new(mockServices.OIDCProvider)

	mockBody := &payload.OauthCallback{
		Code: utils.Ptr("code"),
	}

	mockToken := &oauth2.Token{
		AccessToken:  "accessToken",
		Expiry:       time.Now().Add(24 * time.Hour),
		RefreshToken: "refreshToken",
		TokenType:    "type",
	}
	
	mockOAuth2Client.EXPECT().Exchange(mock.Anything, mock.Anything).Return(mockToken, nil)
	mockOIdcProvider.EXPECT().UserInfo(mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get user info"))

	underTest := services.NewLoginService(mockUserRepo, mockOAuth2Client, mockOIdcProvider)

	// Test Success
	result, err := underTest.OAuthSetup(mockBody)

	is.Nil(result)
	is.NotNil(err)
	is.ErrorIs(err, fmt.Errorf("failed to get user info"))
}

func TestLoginService(t *testing.T) {
	suite.Run(t, new(LoginServiceTestSuite))
}
