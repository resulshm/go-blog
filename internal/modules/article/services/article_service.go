package services

import (
	"errors"

	articleModel "github.com/resulshm/go-blog/internal/modules/article/models"
	articleRepository "github.com/resulshm/go-blog/internal/modules/article/repositories"
	"github.com/resulshm/go-blog/internal/modules/article/requests/articles"
	articleResponse "github.com/resulshm/go-blog/internal/modules/article/responses"
	userResponse "github.com/resulshm/go-blog/internal/modules/user/responses"
)

type ArticleService struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: articleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() articleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return articleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() articleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return articleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (articleResponse.Article, error) {
	var response articleResponse.Article

	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return articleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) Store(request articles.StoreRequest, user userResponse.User) (articleResponse.Article, error) {
	var response articleResponse.Article

	article := articleModel.Article{
		Title:   request.Title,
		Content: request.Content,
		UserID:  user.ID,
	}

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("error in creation of article")
	}

	return articleResponse.ToArticle(newArticle), nil
}
