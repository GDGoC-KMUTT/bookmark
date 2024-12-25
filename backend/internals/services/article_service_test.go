package services_test

import (
	"backend/internals/db/models"
	"backend/internals/services"
	"backend/internals/utils"
	mockRepositories "backend/mocks/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleTestSuit struct {
	suite.Suite
}

func (suite *ArticleTestSuit) TestGetAllArticlesWhenSuccess() {
	is := assert.New(suite.T())

	// Arrange
	mockArticleRepo := new(mockRepositories.ArticleRepository)

	// Mock data
	mockId := utils.Ptr[uint64](1)
	mockTitle := utils.Ptr("testtitle")
	mockHref := utils.Ptr("testhref")

	mockArticleRepo.On("FindAllArticles").Return([]models.Article{
		{
			Id:    mockId,
			Title: mockTitle,
			Href:  mockHref,
		},
	}, nil)

	// Test
	underTest := services.NewArticleService(mockArticleRepo)

	// Test Success
	articles, err := underTest.GetAllArticles()
	is.NoError(err)
	is.Len(articles, 1)
	is.Equal(articles[0].Id, mockId)
	is.Equal(articles[0].Title, mockTitle)
	is.Equal(articles[0].Href, mockHref)
}

func (suite *ArticleTestSuit) TestGetAllArticlesWhenFailed() {
	is := assert.New(suite.T())

	mockArticleRepo := new(mockRepositories.ArticleRepository)
	mockArticleRepo.On("FindAllArticles").Return(nil, fmt.Errorf("article not found"))

	underTest := services.NewArticleService(mockArticleRepo)

	articles, err := underTest.GetAllArticles()

	is.Nil(articles)
	is.NotNil(err)
}

func TestArticleService(t *testing.T) {
	suite.Run(t, new(ArticleTestSuit))
}
