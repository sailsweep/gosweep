package main

import (
	"context"
	configService "entdemo/config"
	myconfig "entdemo/config/types"
	serverTypes "entdemo/server/http"
	svc "entdemo/svc"
)

func main() {
	var (
		ctx           context.Context
		config        myconfig.Config
		err           error
		commonService *svc.ServiceContext
		server        *serverTypes.HttpServer
	)
	ctx = context.Background()
	config, err = configService.NewConfig()
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, "config", config)
	commonService, err = svc.NewServiceContext(ctx)
	if err != nil {
		panic(err)
	}
	server = serverTypes.NewHttpServer(ctx)
	err = server.Run(commonService)
	if err != nil {
		panic(err)
	}
	defer func() {
		commonService.Close()
	}()
}
