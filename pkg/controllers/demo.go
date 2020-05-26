package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyongxi/go-common/server"
	"github.com/shiyongxi/go-gin-web/model/request"
	"github.com/shiyongxi/go-gin-web/pkg/services"
)

type DemoControllers struct {
	server.Controller
}

func NewDemoControllers() *DemoControllers {
	return &DemoControllers{
		server.Controller{},
	}
}

func (ctl *DemoControllers) Test(ctx *gin.Context) {
	var params request.DemoRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctl.ParamsException(ctx, err)
		return
	}
	result, err := services.NewDemoService(ctx.Request.Context()).GetInfo(&params)
	if err != nil {
		ctl.ServiceCodeException(ctx, 400002, err)
		return
	}

	ctl.Response(ctx, result)
}
