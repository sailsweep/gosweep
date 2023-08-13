package http

import (
	"context"
	httpTypes "entdemo/common/types"
	configTypes "entdemo/config/types"
	"entdemo/server/http/route"
	serviceCtx "entdemo/svc"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Port int `json:"port"`
}

type HttpServer struct {
	Config ServerConfig
	Client interface{}
}

func NewHttpServer(ctx context.Context) *HttpServer {
	var (
		server   *HttpServer
		myconfig configTypes.Config
	)
	server = new(HttpServer)
	myconfig = ctx.Value("config").(configTypes.Config)
	server.Config.Port = myconfig.Server.Port
	server.Init()
	return server
}

func (s *HttpServer) Init() {
	s.Client = gin.Default()
	return
}

func (s *HttpServer) Run(svc *serviceCtx.ServiceContext) error {
	var (
		port     string
		apiRoute *route.ApiRoute
	)
	apiRoute = new(route.ApiRoute)
	port = fmt.Sprintf(":%v", s.Config.Port)
	apiRoute.InitRoutes(svc)
	for _, apiGroup := range apiRoute.ApiGroups {
		s.AddRouer(apiGroup.ServiceContext, apiGroup.Method, apiGroup.Path, apiGroup.Handler)
	}
	return s.Client.(*gin.Engine).Run(port)
}

func (s *HttpServer) AddRouer(svc *serviceCtx.ServiceContext, method string, path string, handler httpTypes.HandleFunc) {
	var (
		ginHandler gin.HandlerFunc
	)
	ginHandler = func(c *gin.Context) {
		req := httpTypes.Request{
			Body:  c.Request.Body,
			Query: c.Request.URL.Query(),
		}
		resp, err := handler(svc, req)
		if err != nil {
			c.JSON(resp.Code, resp.Msg)
		} else {
			c.JSON(resp.Code, resp.Data)
		}
	}
	s.Client.(*gin.Engine).Handle(method, path, ginHandler)
}
