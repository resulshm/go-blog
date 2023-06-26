package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userRepositories "github.com/resulshm/go-blog/internal/modules/user/repositories"
	"github.com/resulshm/go-blog/pkg/sessions"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = userRepositories.New()
	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)
		user := userRepo.FindByID(userID)
		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Next()
	}
}
