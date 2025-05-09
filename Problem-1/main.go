package main

import (
	"github.com/ajneelansh/2200290130116/Problem-1/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	token := os.Getenv("ACCESS_TOKEN")
	if token == "" {
		log.Fatalf("ACCESS_TOKEN not set in .env")
	}
	router := gin.Default()

	router.GET("/numbers/:numberid", handlers.GetNumbersHandler)

	router.Run(":9876")
}