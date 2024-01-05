package health

import (
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/internal/module"
	"go-web-server-2-practice/modules/test"
)

type Module struct {
	module.UnimplementedModule
}

func New(app *core.App, tS test.Service) *Module {
	service := &Service{
		tS: tS,
	}
	controller := &Controller{
		s: service,
		app: app,
	}
	m := &Module{}
	m.AddController(controller)
	return m
}