package main

import (
	"log"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/config"
	"github.com/DuvanAlbarracin/movies_apigateway/pkg/profile"
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
