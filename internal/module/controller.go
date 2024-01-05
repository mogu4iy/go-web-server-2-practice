package module

import (
	"github.com/gin-gonic/gin"
	"log"
)

type HttpRequesthandler struct{
	Method string
	Path string
	Handlers []gin.HandlerFunc
}

type Controller interface {
	Init() error
	SetPrefix(string)
	AddRoute(router *gin.RouterGroup, controller *HttpRequesthandler)
	InitRoutes()
	mustEmbedUnimplementedController()
}

type UnimplementedController struct {
	Prefix string
	routes map[*gin.RouterGroup][]*HttpRequesthandler
}

func (c *UnimplementedController) Init() error {
	return nil
}

func (c *UnimplementedController) SetPrefix(prefix string) {
	c.Prefix = prefix
}

func (c *UnimplementedController) AddRoute(router *gin.RouterGroup, handler *HttpRequesthandler){
	if c.routes == nil {
		c.routes = make(map[*gin.RouterGroup][]*HttpRequesthandler)
	}
	c.routes[router] = append(c.routes[router], handler)
}

func (c *UnimplementedController) InitRoutes(){
	log.Println("InitRoutes --> ",c.routes)
	for route, requestHandlers := range c.routes {
		for _, requestHandler := range requestHandlers {
			route.Handle(requestHandler.Method, requestHandler.Path, requestHandler.Handlers...)
		}
	}
}

func (c *UnimplementedController) mustEmbedUnimplementedController(){}