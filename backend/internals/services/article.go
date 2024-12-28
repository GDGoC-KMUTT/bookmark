package services

import "backend/internals/entities/payload"

type ArticleService interface {
	GetAllArticles() ([]payload.Article, error)
}
