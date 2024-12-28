package repositories

import (
	"backend/internals/db/models"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) FindAllArticles() ([]models.Article, error) {
	var articles []models.Article

	result := r.db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}
