package handlers

import (
	"encoding/json"
	"time"

	"github.com/ajneelansh/2200290130116/Problem-1/services"
	"github.com/ajneelansh/2200290130116/Problem-1/storage"
	"github.com/ajneelansh/2200290130116/Problem-1/utils"

	"github.com/gin-gonic/gin"
)

var numberWindow = storage.NewNumberWindow(10)

func GetNumbersHandler(c *gin.Context) {
	numberId := c.Param("numberid")

	prevState := numberWindow.GetNumbers()

	numbers, err := services.FetchNumbers(numberId, 500*time.Millisecond)
	if err != nil {
		numbers = []int{}
	}
	numberWindow.AddNumbers(numbers)

	currState := numberWindow.GetNumbers()

	avg := utils.CalculateAverage(currState)

	response := map[string]interface{}{
		"windowPrevState": prevState,
		"windowCurrState": currState,
		"numbers":         numbers,
		"avg":             avg,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(response)
}