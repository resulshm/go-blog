package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/pkg/html"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home Page ",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page ",
		})
	})
}
