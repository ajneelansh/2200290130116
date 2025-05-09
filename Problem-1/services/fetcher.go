package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"os"
)

type NumbersResponse struct {
	Numbers []int `json:"numbers"`
}

func GetAPIUrl(numberId string) string {
	switch numberId {
	case "p":
		return "http://20.244.56.144/evaluation-service/primes"
	case "f":
		return "http://20.244.56.144/evaluation-service/fibo"
	case "e":
		return "http://20.244.56.144/evaluation-service/even"
	case "r":
		return "http://20.244.56.144/evaluation-service/rand"
	default:
		return ""
	}
}

func FetchNumbers(numberId string, timeout time.Duration) ([]int, error) {
	url := GetAPIUrl(numberId)
	if url == "" {
		return nil, errors.New("invalid numberId")
	}

	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("ACCESS_TOKEN")
	if token == "" {
		return nil, errors.New("access token missing in environment")
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("API Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	var numbersResp NumbersResponse
	err = json.NewDecoder(resp.Body).Decode(&numbersResp)
	if err != nil {
		return nil, err
	}

	fmt.Println("[DEBUG] Numbers received:", numbersResp.Numbers)


	return numbersResp.Numbers, nil

}