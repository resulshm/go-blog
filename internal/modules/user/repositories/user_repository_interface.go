package repositories

import userModels "github.com/resulshm/go-blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user userModels.User) userModels.User
}
