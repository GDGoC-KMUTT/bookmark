package repositories

import "backend/internals/db/models"

type ArticleRepository interface {
	FindAllArticles() ([]models.Article, error)
}
