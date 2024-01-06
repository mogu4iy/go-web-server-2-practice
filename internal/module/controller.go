package module

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HTTPHandlerFunc func (ctx *gin.Context) (interface{}, error)

type HTTPRequesthandler struct{
	Method string
	Path string
	Handlers []HTTPHandlerFunc
}

type HTTPResponse struct{
	success bool
	data interface{}
	error string
}

type httpError struct{
	code int
	err error
}

func (e httpError) Error() string {
	return e.err.Error()
}

type Controller interface {
	Init() error
	SetPrefix(string)
	AddRoute(router *gin.RouterGroup, controller *HTTPRequesthandler)
	InitRoutes()
	ApplyHandler(handler HTTPHandlerFunc) gin.HandlerFunc
	AbortWithStatus(code int)
	mustEmbedUnimplementedController()
}

type UnimplementedController struct {
	Prefix string
	routes map[*gin.RouterGroup][]*HTTPRequesthandler
}

func (*UnimplementedController) Init() error {
	return nil
}

func (c *UnimplementedController) SetPrefix(prefix string) {
	c.Prefix = prefix
}

func (c *UnimplementedController) ErrorWithStatus(code int, err error) error {
	return httpError{code: code, err: err}
}

func (c *UnimplementedController) ApplyHandler(handler HTTPHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := handler(ctx)
		if err != nil {
			response := HTTPResponse{
				success: false,
				error: err.Error(),
			}
			var errAsserted httpError
			ok := errors.As(err, &errAsserted)
			if ok {
				ctx.JSON(errAsserted.code, &response)
			}
			ctx.JSON(http.StatusInternalServerError, &response)
		}
		response := HTTPResponse{
			success: true,
			data: data,
		}
		ctx.JSON(http.StatusOK, &response)
	}
}

func (c *UnimplementedController) AddRoute(router *gin.RouterGroup, handler *HTTPRequesthandler){
	if c.routes == nil {
		c.routes = make(map[*gin.RouterGroup][]*HTTPRequesthandler)
	}
	c.routes[router] = append(c.routes[router], handler)
}

func (c *UnimplementedController) InitRoutes(){
	log.Println("InitRoutes --> ",c.routes)
	for route, requestHandlers := range c.routes {
		for _, requestHandler := range requestHandlers {
			handlers := make([]gin.HandlerFunc, len(requestHandler.Handlers))
			for i, h := range  requestHandler.Handlers{
				handlers[i] = c.ApplyHandler(h)
			}
			route.Handle(requestHandler.Method, requestHandler.Path, handlers...)
		}
	}
}

func (*UnimplementedController) mustEmbedUnimplementedController(){}