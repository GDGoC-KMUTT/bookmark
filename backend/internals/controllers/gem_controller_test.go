package controllers

import (
	"backend/internals/entities/payload"
	"backend/internals/entities/response"
	"backend/internals/services"
	mockServices "backend/mocks/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/mock"
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

	app.Get("/profile/totalgems", controller.GetUserGems)

	return app
}

func TestGetUserGemsWhenSuccess(t *testing.T) {
    is := assert.New(t)

    mockProfileService := new(mockServices.ProfileService)

    app := setupTestGemController(mockProfileService)

    expectedGems := payload.GemTotal{
        Total: 100,
    }

    mockProfileService.EXPECT().GetTotalGems(mock.Anything).Return(&expectedGems, nil)

    req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)  // Correct URL path
    req.Header.Set("Authorization", "Bearer mockToken")  // Add mock JWT token to the header

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

    app := setupTestGemController(mockProfileService)

    mockProfileService.EXPECT().GetTotalGems(mock.Anything).Return(nil, fmt.Errorf("failed to fetch total gems"))

    req := httptest.NewRequest(http.MethodGet, "/profile/totalgems", nil)  // Correct URL path
    req.Header.Set("Authorization", "Bearer mockToken")  // Add mock JWT token to the header

    res, err := app.Test(req)

    var errResponse response.GenericError
    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &errResponse)

    is.Nil(err)
    is.Equal(http.StatusInternalServerError, res.StatusCode)
    is.Equal("failed to fetch total gems", errResponse.Message)
}