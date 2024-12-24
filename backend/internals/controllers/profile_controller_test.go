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
	"testing"
)

type ProfileControllerTestSuit struct {
	suite.Suite
}

func setupTestProfileController(mockProfileService *mockServices.ProfileService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}

	app := fiber.New(fiberConfig)

	// Initialize the controller
	profileController := NewProfileController(mockProfileService)

	// Middleware to simulate JWT Locals
	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{"userId": float64(123)} // Simulate a valid userId claim
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	// Register the route
	app.Get("/profile/info", profileController.ProfileUserInfo)
	return app
}

func (suite *ProfileControllerTestSuit) TestProfileUserInfoWhenSuccess() {
	is := assert.New(suite.T())

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	mockUserId := uint64(123)

	mockProfileService.EXPECT().GetUserInfo(mock.Anything).Return(&payload.Profile{
		Id: utils.Ptr(mockUserId),
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/profile/info", nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[payload.Profile])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(mockUserId, *r.Data.Id)
	is.Equal(http.StatusOK, res.StatusCode)
}
func (suite *ProfileControllerTestSuit) TestProfileUserInfoWhenFailedToGetUserProfile() {
	is := assert.New(suite.T())

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	mockProfileService.EXPECT().GetUserInfo(mock.Anything).Return(nil, fmt.Errorf("get user profile error"))

	req := httptest.NewRequest(http.MethodGet, "/profile/info", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	//'err' will typically be nil for most test cases
	//The 'err' in this context doesn't reflect application logic errors.
	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get user profile", errResponse.Message)
}

func (suite *ProfileControllerTestSuit) TestGetUserGemsWhenSuccess() {
	is := assert.New(suite.T())

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	// mockUserId := uint64(123)
	expectedGems := payload.GemTotal{
		Total: 100,
	}

	mockProfileService.EXPECT().GetTotalGems(mock.Anything).Return(expectedGems, nil)

	req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)
	res, err := app.Test(req)

	var responsePayload response.InfoResponse[payload.GemTotal]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(expectedGems.Total, responsePayload.Data.Total)
}

func (suite *ProfileControllerTestSuit) TestGetUserGemsWhenFailedToFetchTotalGems() {
	is := assert.New(suite.T())

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	// mockUserId := uint64(123)

	mockProfileService.EXPECT().GetTotalGems(mock.Anything).Return(payload.GemTotal{}, fmt.Errorf("failed to fetch total gems"))

	req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to fetch total gems", errResponse.Message)
}

func (suite *ProfileControllerTestSuit) TestGetUserGemsWhenInvalidUserId() {
	is := assert.New(suite.T())

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{}
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusBadRequest, res.StatusCode)
	is.Equal("Invalid userId", errResponse.Message)
}

func TestProfileController(t *testing.T) {
	suite.Run(t, new(ProfileControllerTestSuit))
}
