package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/internal/middlewares"
	userCtrl "github.com/resulshm/go-blog/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", userController.HandleLogout)
	}
}
