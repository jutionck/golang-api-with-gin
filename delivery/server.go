package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-api-with-gin/config"
	"github.com/jutionck/golang-api-with-gin/delivery/controller"
	"github.com/jutionck/golang-api-with-gin/repository"
	"github.com/jutionck/golang-api-with-gin/usecase"
)

type appServer struct {
	productUc     usecase.CreateProductUseCase
	productUcList usecase.ListProductUseCase
	engine        *gin.Engine
	host          string
}

func Server() *appServer {
	r := gin.Default()
	productRepo := repository.NewProductRepository()
	productUc := usecase.NewCreateProductUseCase(productRepo)
	productUcList := usecase.NewListProductUseCase(productRepo)
	c := config.NewConfig()
	host := c.Url
	return &appServer{
		productUc:     productUc,
		productUcList: productUcList,
		engine:        r,
		host:          host,
	}
}

func (a *appServer) initControllers() {
	controller.NewProductController(a.engine, a.productUc, a.productUcList)
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
