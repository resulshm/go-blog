package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/pkg/html"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home Page ",
	})
}
