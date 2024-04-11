package handler

import "mezink-assignment/internal/service"

type IHandler interface {
	IRecordHandler
}

type Handler struct {
	RecordHandler
}

func NewHandler(svc service.IService) IHandler {
	return &Handler{
		RecordHandler{
			Service: svc,
		},
	}
}
