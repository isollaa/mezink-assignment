package internal

import (
	"mezink-assignment/config"
	"mezink-assignment/internal/handler"
	"mezink-assignment/internal/service"
)

func ApiHandler(conf config.Config) handler.IHandler {
	svc := service.NewService(conf)
	return handler.NewHandler(svc)
}
