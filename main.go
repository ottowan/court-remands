package main

import "github.com/gin-gonic/gin"

func main() {
	GetMainEngine().Run(":8080")
}

func GetMainEngine() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
	}
	return r
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
