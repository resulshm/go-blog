package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, key string, value string) {
	session := sessions.Default(c)

	session.Set(key, value)
	session.Save()
}

func Flash(c *gin.Context, key string) string {
	session := sessions.Default(c)
	response := session.Get(key)

	session.Save()
	if response != nil {
		return response.(string)
	}

	return ""
}

func Get(c *gin.Context, key string) string {
	session := sessions.Default(c)
	response := session.Get(key)

	session.Save()

	session.Delete(key)
	session.Save()
	if response != nil {
		return response.(string)
	}

	return ""
}
