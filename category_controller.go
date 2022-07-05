package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	router    *gin.Engine
	ctUseCase CategoryUseCase
}

func (cc *CategoryController) createNewCategory(c *gin.Context) {
	var category *Category
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		category := cc.ctUseCase.Insert(category)
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    category,
		})
	}
}

func (cc *CategoryController) getAllCategory(c *gin.Context) {
	category := cc.ctUseCase.List()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    category,
	})
}

func NewCategoryController(router *gin.Engine, ctUseCase CategoryUseCase) *CategoryController {
	ctController := CategoryController{
		router:    router,
		ctUseCase: ctUseCase,
	}
	router.POST("/category", ctController.createNewCategory)
	router.GET("/category", ctController.getAllCategory)
	return &ctController
}
