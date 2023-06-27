package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/internal/modules/article/requests/articles"
	articleService "github.com/resulshm/go-blog/internal/modules/article/services"
	"github.com/resulshm/go-blog/internal/modules/user/helpers"
	"github.com/resulshm/go-blog/pkg/converters"
	"github.com/resulshm/go-blog/pkg/errors"
	"github.com/resulshm/go-blog/pkg/html"
	"github.com/resulshm/go-blog/pkg/old"
	"github.com/resulshm/go-blog/pkg/sessions"
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

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (controller *Controller) Store(c *gin.Context) {
	var request articles.StoreRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}
	user := helpers.Auth(c)

	article, err := controller.articleService.Store(request, user)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
