package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	articleService "github.com/resulshm/go-blog/internal/modules/article/services"
	"github.com/resulshm/go-blog/pkg/html"
)

type Controller struct {
	articleService articleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: articleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page ",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}
