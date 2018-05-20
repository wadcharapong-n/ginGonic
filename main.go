package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

func main() {
	r := gin.Default()
	r.GET("/ping", ping) //get API
	r.POST("/ping", pingPost)  //post API
	r.GET("/user/:name", func(c *gin.Context) {  //get API have parameter
		name := c.Param("name")
		fmt.Println(name)
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "pong",
			"method": "get",
		})
}

func pingPost(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "pong",
			"method": "post",
		})
}
