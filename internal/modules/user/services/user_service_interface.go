package services

import (
	"github.com/resulshm/go-blog/internal/modules/user/requests/auth"
	userResponses "github.com/resulshm/go-blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (userResponses.User, error)
	CheckUserExists(email string) bool
}
