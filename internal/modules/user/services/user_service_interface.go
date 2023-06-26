package services

import (
	"github.com/resulshm/go-blog/internal/modules/user/requests/auth"
	userResponse "github.com/resulshm/go-blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (userResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (userResponse.User, error)
}
