package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-api-with-gin/config"
	"github.com/jutionck/golang-api-with-gin/delivery/controller"
	"github.com/jutionck/golang-api-with-gin/delivery/middleware"
	"github.com/jutionck/golang-api-with-gin/manager"
	"github.com/rs/zerolog"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *zerolog.Logger
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repoManager := manager.NewRepositoryManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", appConfig.ApiHost, appConfig.ApiPort)
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
		log:            infra.Log(),
	}
}

func (a *appServer) initControllers() {
	a.engine.Use(middleware.LogRequestMiddleware(a.log))
	a.engine.Use(middleware.LogResponseMiddleware(a.log))
	controller.NewProductController(a.engine, a.useCaseManager.CreateProductUseCase(), a.useCaseManager.ListProductUseCase(), a.useCaseManager.FindProductUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
