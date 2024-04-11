package service

import (
	"mezink-assignment/config"
	"mezink-assignment/infra"
)

type IService interface {
	IRecordService
}

type Service struct {
	RecordService
}

func NewService(conf config.Config) IService {
	infra := infra.NewConnection(infra.NewConfig(conf))
	return &Service{
		RecordService{
			Config: conf,
			Infra:  infra,
		},
	}
}
