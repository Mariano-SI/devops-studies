package main

import (
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

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://mercado-livre.com.br",
	}

	for _, url := range urls {
		status, err := CheckServer(url)
		if err != nil {
			fmt.Printf("❌ Erro ao acessar %s: %v\n", url, err)
			continue
		}
		fmt.Printf("✅ %s: Status %d\n", url, status)
	}
}