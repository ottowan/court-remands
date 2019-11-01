package main

import (
	"court-remands/remandsintegration"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	GetMainEngine().Run(":8080")
}

func GetMainEngine() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		v1.POST("/remands", PostRemands)
	}
	return r
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}

func PostRemands(c *gin.Context) {

	courtType, _ := strconv.Atoi(c.PostForm("court_type"))
	penaltyRate, _ := strconv.Atoi(c.PostForm("penalty_rate"))
	fineRate, _ := strconv.Atoi(c.PostForm("fine_rate"))
	remandsTimes, _ := strconv.Atoi(c.PostForm("remands_times"))
	remandsDate, _ := strconv.Atoi(c.PostForm("remands_date"))

	check := remandsintegration.CheckRemandsAPI(courtType, penaltyRate, fineRate, remandsTimes, remandsDate)
	c.JSON(200, gin.H{
		"message": check,
	})

}
