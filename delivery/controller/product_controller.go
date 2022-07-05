package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/usecase"
	"net/http"
)

type ProductController struct {
	router        *gin.Engine
	ucProduct     usecase.CreateProductUseCase
	ucProductList usecase.ListProductUseCase
}

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct *model.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": err.Error(),
		})
	} else {
		err := p.ucProduct.CreateProduct(newProduct)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating Product",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": newProduct,
		})
	}
}

func (p *ProductController) findAllProduct(c *gin.Context) {
	products, err := p.ucProductList.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating Product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": products,
	})
}

func NewProductController(
	router *gin.Engine,
	ucProduct usecase.CreateProductUseCase,
	ucProductList usecase.ListProductUseCase) *ProductController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := ProductController{
		router:        router,
		ucProduct:     ucProduct,
		ucProductList: ucProductList,
	}

	// ini method-method nya
	router.POST("/product", controller.createNewProduct)
	router.GET("/product", controller.findAllProduct)
	return &controller
}
