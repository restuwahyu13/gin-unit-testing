package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Biodata struct {
	Name string `name:json`
	Age  int    `age:json`
}

func UnitTestingExampleGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "successfuly",
		"message": "get request successfuly",
	})
}

func UnitTestingExamplePOST(c *gin.Context) {

	var requestBody Biodata
	c.ShouldBindJSON(&requestBody)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "successfuly",
		"message": "post request successfuly",
		"data":    requestBody,
	})
}

func main() {
	router := SetupRouter()
	router.Run(":3000")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", UnitTestingExampleGET)
	router.POST("/", UnitTestingExamplePOST)
	return router
}
