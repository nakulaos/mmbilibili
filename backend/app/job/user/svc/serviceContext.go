package svc

import "backend/app/job/user/conf"

type ServiceContext struct {
}

func NewServiceContext(c *conf.Config) *ServiceContext {
	return &ServiceContext{}
}
