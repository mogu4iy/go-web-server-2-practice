package test

import (
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/internal/module"
)

type Module struct {
	module.UnimplementedModule
	Service Service
}
func New(app *core.App) *Module {
	service := Service{}
	m := &Module{
		Service: service,
	}
	return m
}