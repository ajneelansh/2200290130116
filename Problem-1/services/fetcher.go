package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var apiMap = map[string]string{
	"p": "primes",
	"f": "fibo",
	"e": "even",
	"r": "rand",
}

type NumberResponse struct {
	Numbers []int `json:"numbers"`
}

func FetchNumbers(id string, timeout time.Duration) ([]int, error) {
	api, ok := apiMap[id]
	if !ok {
		return nil, errors.New("invalid number id")
	}

	url := fmt.Sprintf("http://20.244.56.144/evaluation-service/%s", api)

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err 
	}
	defer resp.Body.Close()

	var result NumberResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err 
	}

	return result.Numbers, nil
}