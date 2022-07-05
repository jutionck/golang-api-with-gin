package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	ctUseCase CategoryUseCase
	engine    *gin.Engine
	host      string
}

func (a *appServer) initHandlers() {
	NewCategoryController(a.engine, a.ctUseCase)
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()
	ctRepo := NewCategoryRepository()
	ctUseCase := NewCategoryUseCase(ctRepo)
	host := fmt.Sprintf("%s:%s", "localhost", "8888")
	return &appServer{
		ctUseCase: ctUseCase,
		engine:    r,
		host:      host,
	}
}
