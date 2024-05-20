package main

import (
	"log"

	"github.com/DuvanAlbarracin/movies_api_gateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_api_gateway/pkg/profile"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	r := gin.Default()
	_ = *profile.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
