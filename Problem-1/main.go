package main

import (
	"github.com/ajneelansh/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/numbers/:numberid", handlers.GetNumbersHandler)

	router.Run(":9876")
}