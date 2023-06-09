package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resulshm/go-blog/pkg/config"
	"github.com/resulshm/go-blog/pkg/routing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in config.yml file",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config.Set()

	routing.Init()
	router := routing.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.GetString("App.Name"),
		})
	})

	routing.Serve()
}
