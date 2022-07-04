package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	routerEngine := gin.Default()
	routerEngine.Use(AuthMiddleware()) // implement middleware (global)
	/*
	* grouping: auth
	* www.enigmacamp.com/auth/login
	* www.enigmacamp.com/auth/register
	* www.enigmacamp.com/auth/logout
	*
	* grouping: master
	* www.enigmacamp.com/master/users
	* www.enigmacamp.com/master/students

	* www.enigmacamp.com/api/auth/login
	 */

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})
	rgAuth := routerEngine.Group("/api/auth")
	rgAuth.POST("/login", login)

	rgMaster := routerEngine.Group("/api/master")
	rgMaster.GET("/greeting/:name", greeting)

	err := routerEngine.Run()
	if err != nil {
		panic(err)
	} // secara default menggunakan port :8080

}

//func greeting(c *gin.Context) {
//	name := c.Param("name") // with parameter :name at path /greeting
//	c.String(http.StatusOK, "Hello....%s", name)
//}

// Query Param / String in Path -> ?key=value
func greeting(c *gin.Context) {
	name := c.Param("name")    // with parameter :name at path /greeting
	kec := c.Query("username") // Query Param / String in Path -> ?key=value
	kel := c.Query("password") // Query Param / String in Path -> ?key=value
	c.String(http.StatusOK, "Hello %s saat ini kamu berada kec %s kel %s", name, kec, kel)
}

//func login(c *gin.Context) {
//	username := c.PostForm("username")
//	c.PostForm("password")
//	c.String(http.StatusOK, "Hello %s", username)
//}

/*
* Model Binding & Validation
* -> ShouldBind, ShouldBindJSON -> using struct tag binding:required
 */

func login(c *gin.Context) {
	var uc UserCredential
	if err := c.BindJSON(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": "ini_token",
		})
	}
}
