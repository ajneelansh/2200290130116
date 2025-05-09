package main

import (
	"encoding/json"
    "fmt"
    "log"
    "math"
    "net/http"
    "strconv"
    "sync"
    "time"

    "github.com/ajneelansh/handlers"
    "github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

    router.GET("/stocks/:ticker", handlers.GetAverageStockPrice())
    router.GET("/stockcorrelation", handlers.GetStockCorrelation())

    log.Println("Server running on :8080")
    router.Run(":8080")

	
}