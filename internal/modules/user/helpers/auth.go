package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	userRepositories "github.com/resulshm/go-blog/internal/modules/user/repositories"
	userResponse "github.com/resulshm/go-blog/internal/modules/user/responses"
	"github.com/resulshm/go-blog/pkg/sessions"
)

func Auth(c *gin.Context) userResponse.User {
	var response userResponse.User
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	userRepo := userRepositories.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}
	return userResponse.ToUser(user)
}
