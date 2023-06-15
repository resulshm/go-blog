package routes

import (
	"github.com/gin-gonic/gin"
	articleRoutes "github.com/resulshm/go-blog/internal/modules/article/routes"
	homeRoutes "github.com/resulshm/go-blog/internal/modules/home/routes"
	userRoutes "github.com/resulshm/go-blog/internal/modules/user/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
