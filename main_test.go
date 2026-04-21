package main

import "testing"

func TestCheckServer(t *testing.T) {
	url := "https://google.com"
	expectedStatus := 200

	status, err := CheckServer(url)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if status != expectedStatus {
		t.Errorf("Esperado %d, mas obteve %d", expectedStatus, status)
	}
}