package main

import "testing"

func TestCheckHealth(t *testing.T) {
	// Teste com uma URL que sabemos que deve estar online
	url := "https://google.com"
	if !CheckHealth(url) {
		t.Errorf("Esperava que %s estivesse ONLINE, mas retornou OFFLINE", url)
	}

	// Teste com uma URL inválida
	urlInvalida := "https://url-que-nao-existe-123.com"
	if CheckHealth(urlInvalida) {
		t.Errorf("Esperava que %s estivesse OFFLINE, mas retornou ONLINE", urlInvalida)
	}
}
