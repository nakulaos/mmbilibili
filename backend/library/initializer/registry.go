package initializer

import (
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"log"
)

var Registry registry.Registry

type RegistryInitializer struct {
	RegistryAddr string
}

func NewRegistryInitializer(registryAddr string) *RegistryInitializer {
	return &RegistryInitializer{RegistryAddr: registryAddr}
}

func (m *RegistryInitializer) Init() error {
	config := consulapi.DefaultConfig()
	config.Address = m.RegistryAddr
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Registry = consul.NewConsulRegister(consulClient)
	return nil
}

func InitHertzRegistry(addr string) registry.Registry {
	m := NewRegistryInitializer(addr)
	if err := m.Init(); err != nil {
		log.Fatal(err)
	}
	return Registry
}
