package services

import (
	"github.com/resulshm/go-blog/internal/modules/article/requests/articles"
	articleResponse "github.com/resulshm/go-blog/internal/modules/article/responses"
	userResponse "github.com/resulshm/go-blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() articleResponse.Articles
	GetStoriesArticles() articleResponse.Articles
	Find(id int) (articleResponse.Article, error)
	Store(request articles.StoreRequest, user userResponse.User) (articleResponse.Article, error)
}
