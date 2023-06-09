package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/internal/providers/routes"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
