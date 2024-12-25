package controllers

import (
	"backend/internals/entities/response"
	"backend/internals/services"

	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	articleSvc services.ArticleService
}

func NewArticleController(articleSvc services.ArticleService) ArticleController {
	return ArticleController{
		articleSvc: articleSvc,
	}
}

// GetAllArticles
// @ID getAllArticles
// @Tags article
// @Summary Get all articles
// @Accept json
// @Produce json
// @Success 200 {object} response.InfoResponse[[]payload.Article]
// @Failure 400 {object} response.GenericError
// @Router /article [get]
func (r *ArticleController) GetAllArticles(c *fiber.Ctx) error {

	articles, err := r.articleSvc.GetAllArticles()
	if err != nil {
		return &response.GenericError{Err: err, Message: "failed to get all articles"}
	}

	return response.Ok(c, articles)
}
