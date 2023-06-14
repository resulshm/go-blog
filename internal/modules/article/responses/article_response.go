package responses

import (
	articleModels "github.com/resulshm/go-blog/internal/modules/article/models"
	userResponse "github.com/resulshm/go-blog/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      userResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModels.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt.Format("2006/01/02"),
		Image:     "/assets/img/demopic/10.jpg",
		User:      userResponse.ToUser(article.User),
	}
}

func ToArticles(articles []articleModels.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
