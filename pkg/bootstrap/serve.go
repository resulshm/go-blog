package bootstrap

import (
	"github.com/resulshm/go-blog/pkg/config"
	"github.com/resulshm/go-blog/pkg/database"
	"github.com/resulshm/go-blog/pkg/html"
	"github.com/resulshm/go-blog/pkg/routing"
	"github.com/resulshm/go-blog/pkg/sessions"
	"github.com/resulshm/go-blog/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	routing.RegisterRoutes()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.Serve()
}
