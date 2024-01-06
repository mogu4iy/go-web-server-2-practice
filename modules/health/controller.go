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
	c.AddRoute(hRouter, &module.HTTPRequesthandler{
		Method: http.MethodGet,
		Path: "/",
		Handlers: []module.HTTPHandlerFunc{c.Status},
	})
	hV1Router := c.app.V1Router.Group(c.Prefix)
	c.AddRoute(hV1Router, &module.HTTPRequesthandler{
		Method: http.MethodGet,
		Path: "/",
		Handlers: []module.HTTPHandlerFunc{c.Status},
		})
	return nil	
}

func (c *Controller) Status(ctx *gin.Context) (interface{}, error) {
	return c.s.status(), nil
}