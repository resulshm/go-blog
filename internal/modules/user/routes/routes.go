package routes

import (
	"github.com/gin-gonic/gin"
	userCtrl "github.com/resulshm/go-blog/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()
	router.GET("/register", userController.Register)
	router.POST("/register", userController.HandleRegister)

	router.GET("/login", userController.Login)
	router.POST("/login", userController.HandleRegister)
}
