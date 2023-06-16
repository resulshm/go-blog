package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/internal/modules/user/requests/auth"
	userService "github.com/resulshm/go-blog/internal/modules/user/services"
	"github.com/resulshm/go-blog/pkg/html"
)

type Controller struct {
	userService userService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: userService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.userService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	log.Printf("The user created successfully with name %s", user.Name)
	c.Redirect(http.StatusFound, "/")
}