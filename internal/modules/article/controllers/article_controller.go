package controllers

import (
	"net/http"
	"strconv"

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

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{
			"title":   "Server error",
			"message": "error in converting id",
		})
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"title":   "Article is not found",
			"message": err.Error(),
		})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{
		"title":   "Show article",
		"article": article,
	})
}
