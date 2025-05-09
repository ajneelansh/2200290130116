package handlers

import(
	"encoding/json"
    "fmt"
    "log"
    "math"
    "net/http"
    "strconv"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
)

func GetAverageStockPrice() gin.HandlerFunc{
	return func(c *gin.Context){
		ticker := c.Param("ticker")
		minutesStr := c.DefaultQuery("minutes", "30")
		minutes, err := strconv.Atoi(minutesStr)
		if err != nil || minutes <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minutes"})
			return
		}
	
		aggregation := c.DefaultQuery("aggregation", "average")
	
		prices := fetchPrices(ticker, minutes)
	
		if len(prices) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No prices found"})
			return
		}
	
		var sum float64
		for _, p := range prices {
			sum += p.Price
		}
		avg := sum / float64(len(prices))
	
		c.JSON(http.StatusOK, gin.H{
			"averageStockPrice": avg,
			"priceHistory":      prices,
		})
	}
}

func fetchPrices(ticker string, minutes int) []StockPriceEntry {
    url := fmt.Sprintf("%s/%s?minutes=%d", baseAPI, ticker, minutes)
    resp, err := http.Get(url)
    if err != nil {
        log.Println("API error:", err)
        return []StockPriceEntry{}
    }
    defer resp.Body.Close()

    var result []map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&result)
    if err != nil {
        log.Println("Decode error:", err)
        return []StockPriceEntry{}
    }

    var fetched []StockPriceEntry
    for _, item := range result {
        t, err := time.Parse(time.RFC3339Nano, item["lastUpdatedAt"].(string))
        if err != nil {
            continue
        }
        fetched = append(fetched, StockPriceEntry{
            Price:         item["price"].(float64),
            LastUpdatedAt: t,
        })
    }

    return fetched
}
