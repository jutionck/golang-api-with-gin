package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	var db []Category
	categoryRepo := NewCategoryRepository(db)
	categoryUseCase := NewCategoryUseCase(categoryRepo)

	rg := r.Group("/api")
	rg.POST("/category", func(c *gin.Context) {
		var category Category
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			category := categoryUseCase.Insert(&category)
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    category,
			})
		}
	})

	rg.GET("/category", func(c *gin.Context) {
		categories := categoryUseCase.List()
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    categories,
		})
	})
	err := r.Run()
	if err != nil {
		panic(err)
	} // secara default menggunakan port :8080

}
