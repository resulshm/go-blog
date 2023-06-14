package responses

import (
	"fmt"

	userModel "github.com/resulshm/go-blog/internal/modules/user/models"
)

type User struct {
	ID    uint
	Name  string
	Email string
	Image string
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
