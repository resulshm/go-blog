package services

import (
	"errors"

	userModel "github.com/resulshm/go-blog/internal/modules/user/models"
	userRepository "github.com/resulshm/go-blog/internal/modules/user/repositories"
	"github.com/resulshm/go-blog/internal/modules/user/requests/auth"
	userResponse "github.com/resulshm/go-blog/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository userRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: userRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (userResponse.User, error) {
	var response userResponse.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error in hashing password")
	}
	user := userModel.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	newUser := userService.userRepository.Create(user)
	if newUser.ID == 0 {
		return response, errors.New("error in creating new user")
	}

	return userResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}

	return false
}
