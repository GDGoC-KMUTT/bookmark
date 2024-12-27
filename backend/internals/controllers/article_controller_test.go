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
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ArticleControllerTestSuit struct {
	suite.Suite
}

func setupTestArticleController(mockArticleService *mockServices.ArticleService) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

	articleController := NewArticleController(mockArticleService)

	app.Use(func(c *fiber.Ctx) error {
		token := &jwt.Token{}
		claims := jwt.MapClaims{"userId": float64(123)} // Simulate a valid userId claim
		token.Claims = claims
		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/article", articleController.GetAllArticles)

	return app
}

func (suite *ArticleControllerTestSuit) TestGetAllArticlesWhenSuccess() {
	is := assert.New(suite.T())

	mockArticleService := new(mockServices.ArticleService)

	app := setupTestArticleController(mockArticleService)

	mockId := utils.Ptr(uint64(1))
	mockTitle := utils.Ptr("testtitle")
	mockHref := utils.Ptr("testhref")

	mockArticleService.EXPECT().GetAllArticles().Return([]payload.Article{{
		Id:    mockId,
		Title: mockTitle,
		Href:  mockHref,
	}}, nil)

	req := httptest.NewRequest(http.MethodGet, "/article", nil)
	res, err := app.Test(req)

	r := new(response.InfoResponse[[]payload.Article])
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &r)

	is.Nil(err)
	is.Equal(mockId, r.Data[0].Id)
	is.Equal(mockTitle, r.Data[0].Title)
	is.Equal(mockHref, r.Data[0].Href)
	is.Equal(http.StatusOK, res.StatusCode)

}

func (suite *ArticleControllerTestSuit) TestGetAllArticlesWhenFailedToFetchAllArticles() {
	is := assert.New(suite.T())

	mockArticleService := new(mockServices.ArticleService)

	app := setupTestArticleController(mockArticleService)

	mockArticleService.EXPECT().GetAllArticles().Return(nil, fmt.Errorf("get all articles error"))

	req := httptest.NewRequest(http.MethodGet, "/article", nil)
	res, err := app.Test(req)

	var errResponse response.GenericError
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &errResponse)

	is.Nil(err)
	is.Equal(http.StatusInternalServerError, res.StatusCode)
	is.Equal("failed to get all articles", errResponse.Message)

}

func TestArticleController(t *testing.T) {
	suite.Run(t, new(ArticleControllerTestSuit))
}
