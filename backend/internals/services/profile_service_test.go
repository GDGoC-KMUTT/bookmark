package services_test

import (
	"backend/internals/config"
	"backend/internals/db/models"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type ProfileTestSuit struct {
	suite.Suite
}

func (suite *ProfileTestSuit) TestGetUserInfoWhenSuccess() {
	is := assert.New(suite.T())
	// Arrange
	mockUserRepo := new(mockRepositories.UserRepository)

	mockUserId := utils.Ptr[uint64](1)
	mockFirstName := utils.Ptr("testfn")
	mockLastName := utils.Ptr("testln")
	mockEmail := utils.Ptr("test@gmail.com")
	mockPhotoUrl := utils.Ptr("url")

	mockUserRepo.EXPECT().FindUserByID(mockUserId).Return(&models.User{
		Id:        mockUserId,
		Firstname: mockFirstName,
		Lastname:  mockLastName,
		Email:     mockEmail,
		PhotoUrl:  mockPhotoUrl,
	}, nil)

	underTest := services.NewLoginService(mockUserRepo)

	// Test Success
	convList, err := underTest.GetUserInfo(strconv.Itoa(int(*mockUserId)))

	is.Equal(&mockTopicId, convList[0].ID)
	is.Equal("test1", *convList[0].Topic)
	is.NoError(err)
}

func TestProfileService(t *testing.T) {
	config.BootConfiguration()
	suite.Run(t, new(ProfileTestSuit))
}
