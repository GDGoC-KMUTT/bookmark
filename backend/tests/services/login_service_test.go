package services_test

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LoginServiceTestSuite struct {
	suite.Suite
}

func (suite *LoginServiceTestSuite) TestGetOrCreateUserFromClaimWhenSuccess() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)

	// Mock data
	mockUserInfo := &oidc.UserInfo{
		Email:         "test@gmail.com",
		Profile:       "profile",
		EmailVerified: true,
		Subject:       "test subject",
	}

	mockUser := &models.User{
		Id:        utils.Ptr[uint64](1),
		Email:     utils.Ptr("test@email.com"),
		Firstname: utils.Ptr("fn"),
		Lastname:  utils.Ptr("ln"),
		PhotoUrl:  nil,
	}

	// mock user repo
	mockUserRepo.EXPECT().FindFirstUserByOid(mock.Anything).Return(mockUser, nil)

	// Test
	underTest := services.NewLoginService(mockUserRepo)

	// Test Success
	_, err := underTest.GetOrCreateUserFromClaims(mockUserInfo)

	//is.NotNil(userInfo)
	//is.Equal(*mockUser.Id, *userInfo.Id)
	is.NoError(err)
}

func TestLoginService(t *testing.T) {
	config.BootConfiguration()
	suite.Run(t, new(LoginServiceTestSuite))
}
