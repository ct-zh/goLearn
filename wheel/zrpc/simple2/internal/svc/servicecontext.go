package svc

import "github.com/ct-zh/goLearn/wheel/zrpc/simple2/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
