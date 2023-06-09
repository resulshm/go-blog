package bootstrap

import (
	"github.com/resulshm/go-blog/pkg/config"
	"github.com/resulshm/go-blog/pkg/html"
	"github.com/resulshm/go-blog/pkg/routing"
	"github.com/resulshm/go-blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.Serve()
}
