package core

import (
	"github.com/gin-gonic/gin"
	"go-web-server-2-practice/internal/module"
	"log"
)

type App struct {
	Engine *gin.Engine
	Router *gin.RouterGroup
	V1Router *gin.RouterGroup
	VRouter *gin.RouterGroup
	modules []module.Module
}

func (a *App) AddModule(m module.Module) {
	a.modules = append(a.modules, m)
}

func (a *App) Init() error{
	log.Println("Init App --> ", a.modules)
	engine := gin.Default()
	a.VRouter = engine.Group("v2")
	a.V1Router = engine.Group("v1")
	a.Router = engine.Group("/")
	for _, m := range a.modules {
		err:= m.Init()
		if err != nil{
			return err
		}
	}
	return nil
}