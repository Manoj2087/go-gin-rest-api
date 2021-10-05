package main

import (
	"github.com/gin-gonic/gin"

	"gin-rest-api/handlers"
	"gin-rest-api/models"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.GET("/status", handlers.StatusHandler)

	router.GET("/books", handlers.GetAllBooks)
	router.GET("/book/:id", handlers.GetBook)
	router.POST("/book", handlers.AddBook)
	router.PATCH("/book/:id", handlers.UpdateBook)
	router.DELETE("/book/:id", handlers.DeleteBook)

	router.Run(":4000")
}
