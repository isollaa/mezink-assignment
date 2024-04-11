package main

import (
	"fmt"
	"mezink-assignment/config"
	"mezink-assignment/internal"
	"mezink-assignment/transport/route"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	conf := config.Get()
	route.Endpoints(&route.Route{
		Handler: internal.ApiHandler(conf),
		Engine:  engine,
	})

	engine.Run(fmt.Sprintf(":%d", conf.ServerPort))
}
