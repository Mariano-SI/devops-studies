package main

import (
	"fmt"
	"net/http"
	"time"
)

// CheckHealth recebe uma URL e retorna true se o status for 200
func CheckHealth(url string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://mercado-livre.com.br",
	}

	for _, url := range urls {
		status := "OFFLINE"
		if CheckHealth(url) {
			status = "ONLINE"
		}
		fmt.Printf("[%s] - %s\n", status, url)
	}
}
