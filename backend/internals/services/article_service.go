package services

import (
	"backend/internals/entities/payload"
	"backend/internals/repositories"
)

type articleService struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

func (r *articleService) GetAllArticles() ([]payload.Article, error) {
	articles, tx := r.articleRepo.FindAllArticles()
	if tx != nil {
		return nil, tx
	}

	var result []payload.Article
	for _, article := range articles {
		result = append(result, payload.Article{
			Id:    article.Id,
			Title: article.Title,
			Href:  article.Href,
		})
	}
	return result, nil

}
