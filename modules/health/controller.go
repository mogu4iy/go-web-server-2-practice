package health

import (
	"github.com/gin-gonic/gin"
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/internal/module"
	"net/http"
)

type Controller struct {
	module.UnimplementedController
	app *core.App
	s *Service
}

func (c *Controller) Init() error {
	c.SetPrefix("health")
	hRouter := c.app.Router.Group(c.Prefix)
	c.AddRoute(hRouter, &module.HttpRequesthandler{
		Method: http.MethodGet,
		Path: "/",
		Handlers: []gin.HandlerFunc{c.Status},
	})
	hV1Router := c.app.V1Router.Group(c.Prefix)
	c.AddRoute(hV1Router, &module.HttpRequesthandler{
		Method: http.MethodGet,
		Path: "/",
		Handlers: []gin.HandlerFunc{c.Status},
		})
	return nil	
}

func (c *Controller) Status(g *gin.Context) {
	c.s.status()
}