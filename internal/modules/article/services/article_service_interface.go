package services

import articleModels "github.com/resulshm/go-blog/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []articleModels.Article
	GetStoriesArticles() []articleModels.Article
}
