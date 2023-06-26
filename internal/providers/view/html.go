package view

import (
	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/internal/modules/user/helpers"
	"github.com/resulshm/go-blog/pkg/converters"
	"github.com/resulshm/go-blog/pkg/sessions"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(c, "old"))
	data["AUTH"] = helpers.Auth(c)
	return data
}
