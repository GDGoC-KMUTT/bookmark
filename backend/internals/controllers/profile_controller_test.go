package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type ProfileControllerTestSuit struct {
	suite.Suite
}

func (suite *ProfileControllerTestSuit) TestProfileUserInfo() {
	is := assert.New(suite.T())
	// Setup the app as it is done in the main function
	app := fiber.New(fiber.Config{})
	//
	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Perform the request using app.Test
	res, err := app.Test(req, -1)

	// Verify that no error occurred
	is.Nil(err)

	// Verify the status code
	is.Equal(200, res.StatusCode)

	// Read the response body
	body, _ := io.ReadAll(res.Body)
	//
	//// Verify the response body
	is.Equal("OK", string(body))

	//mockProfileService := new(mockServices.ProfileService)
	//
	//underTest := NewProfileController(mockProfileService)
	//
	//response := underTest.ProfileUserInfo()
}

func TestProfileController(t *testing.T) {
	suite.Run(t, new(ProfileControllerTestSuit))
}
