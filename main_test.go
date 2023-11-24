package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPriceHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getPriceHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/price")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGetData(t *testing.T) {
	mockResponse := `{"id": "bitcoin","symbol": "BTC","name": "Bitcoin","priceUsd": "50000","changePercent24Hr": "2.5"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/price")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
