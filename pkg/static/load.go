package static

import "github.com/gin-gonic/gin"

func LoadStatic(router *gin.Engine) {
	router.Static("/assets", "./assets")
}
