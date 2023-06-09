package routing

import (
	"fmt"
	"log"

	"github.com/resulshm/go-blog/pkg/config"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		log.Fatal("Error in routing")
	}
}
