package user

import (
	"github.com/gin-gonic/gin"
	"go-web-server-2-practice/core"
	"go-web-server-2-practice/internal/module"
	dbmodels "go-web-server-2-practice/models"
	"net/http"
)

type Controller struct {
	module.UnimplementedController
	app *core.App
	db *core.DB
}

func (c *Controller) Init() error {
	c.SetPrefix("user")
	uVRouter := c.app.VRouter.Group(c.Prefix)
	c.AddRoute(uVRouter, &module.HTTPRequesthandler{
		Method: http.MethodGet,
		Path: "/",
		Handlers: []module.HTTPHandlerFunc{},
	})
	return nil
}

func (c *Controller) Create(ctx *gin.Context) (interface{}, error) {
	data := &CreateRequestDto{}
	err := ctx.ShouldBindJSON(data)
	if err != nil {
		return nil, c.ErrorWithStatus(http.StatusBadRequest, err)
	}
	user := &dbmodels.User{
		Name: data.Name,
		Email: data.Email,
		PasswordHash: data.Password,
	}	
	result := c.db.Engine.Create(user)
	if result.Error != nil	{
		return nil, c.ErrorWithStatus(http.StatusBadRequest, result.Error)
	}
	return &CreateResponseDto{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}