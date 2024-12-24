package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	mockServices "backend/mocks/services"
	"github.com/stretchr/testify/mock"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupTestGemController(profileSvc services.ProfileService) *fiber.App {
	app := fiber.New()

	controller := NewProfileController(profileSvc)

	app.Use(func(c *fiber.Ctx) error {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["userId"] = float64(123)
		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/api/profile/totalgems", controller.GetUserGems)

	return app
}

func TestGetUserGemsWhenSuccess(t *testing.T) {
	is := assert.New(t)

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	expectedGems := payload.GemTotal{
		Total: 100,
	}

	mockProfileService.EXPECT().GetTotalGems(mock.Anything).Return(&expectedGems, nil)

	req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)
	res, err := app.Test(req)

	var responsePayload response.InfoResponse[payload.GemTotal]
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &responsePayload)

	is.Nil(err)
	is.Equal(http.StatusOK, res.StatusCode)
	is.Equal(expectedGems.Total, responsePayload.Data.Total)
}

func TestGetUserGemsWhenFailedToFetchTotalGems(t *testing.T) {
	is := assert.New(t)

	mockProfileService := new(mockServices.ProfileService)

	app := setupTestProfileController(mockProfileService)

	mockProfileService.EXPECT().GetTotalGems(mock.Anything).Return(&payload.GemTotal{}, fmt.Errorf("failed to fetch total gems"))

	req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to fetch total gems", errResponse.Message)
}

func TestGetUserGemsWhenInvalidUserId(t *testing.T) {
	is := assert.New(t)

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