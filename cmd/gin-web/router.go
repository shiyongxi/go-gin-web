package main

import (
	"github.com/shiyongxi/go-gin-web/pkg/controllers"
)

func (svc *Server) Router() {
	var (
		demoCtl = controllers.NewDemoControllers()
	)

	authRouter := svc.BaseServer.RouterGroup.Group("/v1")
	droneRouter := authRouter.Group("/demo")
	{
		droneRouter.POST("/test", demoCtl.Test)
	}
}

// grpc server registered
//func (svc *Server) RegisteredRpcServer(grpcSvc *grpc.Server) {
//	user.RegisterUserServer(grpcSvc, userPB.NewUserInfoGrpc(svc.Logger))
//}
