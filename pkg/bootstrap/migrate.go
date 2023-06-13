package bootstrap

import (
	"github.com/resulshm/go-blog/internal/database/migration"
	"github.com/resulshm/go-blog/pkg/config"
	"github.com/resulshm/go-blog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
