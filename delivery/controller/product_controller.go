package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-api-with-gin/delivery/api"
	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/usecase"
	"github.com/jutionck/golang-api-with-gin/utils"
)

type ProductController struct {
	router        *gin.Engine
	ucProduct     usecase.CreateProductUseCase
	ucProductList usecase.ListProductUseCase
	api.BaseApi
}

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct model.Product
	err := p.ParseRequestBody(c, &newProduct)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucProduct.CreateProduct(&newProduct)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newProduct)
}

func (p *ProductController) findAllProduct(c *gin.Context) {
	products, err := p.ucProductList.List()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, products)
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
