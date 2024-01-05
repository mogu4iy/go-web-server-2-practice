package module

import (
	"log"
	_ "log"
)

type Module interface {
	Init() error
	AddController(Controller)
	InitControllers () error
	mustEmbedUnimplementedModule()
}

type UnimplementedModule struct {
	controllers []Controller
}

func (m *UnimplementedModule) Init() error{
	log.Println("Init Module --> ", m)
	err := m.InitControllers()
	if err != nil {
		return err
	}
	return nil
}

func (m *UnimplementedModule) AddController(s Controller) {
	m.controllers = append(m.controllers, s)
}

func (m *UnimplementedModule) InitControllers() error {
	log.Println("InitControllers --> ", m.controllers)
	for _, c := range m.controllers {
		err := c.Init()
		if err != nil {
			return err
		}
		c.InitRoutes()
	}
	return nil
}

func (m *UnimplementedModule) mustEmbedUnimplementedModule() {}
