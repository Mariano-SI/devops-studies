package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CheckServer faz o GET em uma URL e retorna o status
func CheckServer(url string) (int, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

type HealthResult struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://mercado-livre.com.br",
	}

	var results []HealthResult
	for _, url := range urls {
		status, err := CheckServer(url)
		result := HealthResult{URL: url, Status: status}
		if err != nil {
			result.Error = err.Error()
			result.Status = 0
		}
		results = append(results, result)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	fmt.Println("🚀 Health Checker web server iniciado na porta 8080!")
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}
