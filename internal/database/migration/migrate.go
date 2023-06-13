package migration

import (
	"fmt"
	"log"

	articleModels "github.com/resulshm/go-blog/internal/modules/article/models"
	userModels "github.com/resulshm/go-blog/internal/modules/user/models"
	"github.com/resulshm/go-blog/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Fatal("Can't migrate database...")
	}

	fmt.Println("Migration done...")
}
