package user

import (
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/internal/module"
)

type Module struct {
	module.UnimplementedModule
}

func New(app *core.App, db *core.DB) *Module {
	controller := &Controller{
		app: app,
		db: db,
	}
	m := &Module{}
	m.AddController(controller)
	return m
}