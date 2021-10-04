package main

import (
	"github.com/gin-gonic/gin"

	"gin-rest-api/models"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	statusHandler := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	}

	router.GET("/status", statusHandler)
	router.Run(":4000")
}
