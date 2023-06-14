package services

import (
	articleModels "github.com/resulshm/go-blog/internal/modules/article/models"
	articleRepository "github.com/resulshm/go-blog/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: articleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() []articleModels.Article {
	return articleService.articleRepository.List(4)
}

func (articleService *ArticleService) GetStoriesArticles() []articleModels.Article {
	return articleService.articleRepository.List(6)
}
