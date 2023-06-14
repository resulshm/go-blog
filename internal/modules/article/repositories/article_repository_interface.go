package repositories

import (
	articleModels "github.com/resulshm/go-blog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []articleModels.Article
}