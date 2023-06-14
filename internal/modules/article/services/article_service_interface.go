package services

import (
	articleResponse "github.com/resulshm/go-blog/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() articleResponse.Articles
	GetStoriesArticles() articleResponse.Articles
}
