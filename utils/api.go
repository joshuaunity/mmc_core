package utils

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"	
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling log entries: %v", err)
		return
	}
	w.Write(res)
}

func PaginationParams(r *http.Request) (int, int) {
	// Get the 'page' query param
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default to page 1 if not provided or invalid
	}

	// Get the 'limit' query param
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // Default limit to 10 if not provided or invalid
	}

	return page, limit
}

func APIServiceHealthCheck(url string) error {
	// call health check api
	response, err := http.Get(url + "/v1/health")
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// Check for non-200 HTTP status codes
	if response.StatusCode != http.StatusOK {
		return errors.New("API Service is not healthy")
	}

	return nil
}

func CheckHealthWithRetry(url string, maxRetries int, initialDelay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = APIServiceHealthCheck(url)
		if err == nil {
			return nil
		}

		// Exponential backoff:
		delay := time.Duration(math.Pow(2, float64(i))) * initialDelay
		time.Sleep(delay)
	}

	return err
}
