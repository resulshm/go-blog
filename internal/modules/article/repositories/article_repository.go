package repositories

import (
	articleModels "github.com/resulshm/go-blog/internal/modules/article/models"
	"github.com/resulshm/go-blog/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []articleModels.Article {
	var articles []articleModels.Article

	articleRepository.DB.Limit(limit).Joins("User").Find(&articles)

	return articles
}
