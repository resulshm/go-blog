package bootstrap

import (
	"github.com/resulshm/go-blog/pkg/config"
	"github.com/resulshm/go-blog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
