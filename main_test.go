package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetEndpointsFromURL(t *testing.T) {
	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<a href="http://example.com/page1">Page 1</a>
			<a href="http://example.com/page2">Page 2</a>
			<link rel="stylesheet" href="http://example.com/style.css">
			<img src="http://example.com/image.jpg">
		`))
	}))
	defer mockServer.Close()

	url := mockServer.URL
	endpoints, err := getEndpointsFromURL(url)
	if err != nil {
		t.Errorf("getEndpointsFromURL(%s) returned an error: %v", url, err)
	}

	expectedEndpoints := []string{"http://example.com/page1", "http://example.com/page2"}
	for i, expected := range expectedEndpoints {
		if endpoints[i] != expected {
			t.Errorf("Expected endpoint %s at index %d, got %s", expected, i, endpoints[i])
		}
	}
}

func TestCheckLatency(t *testing.T) {
	// Mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Simulate latency
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	url := mockServer.URL
	timeout := 5 * time.Second // Increase timeout duration
	latency, err := checkLatency(url, timeout)
	if err != nil {
		t.Errorf("checkLatency(%s, %v) returned an error: %v", url, timeout, err)
	}

	expected := time.Duration(2 * time.Second)
	// Allow for a small difference in latency (e.g., 100 milliseconds)
	maxDeviation := 100 * time.Millisecond
	if latency < expected-maxDeviation || latency > expected+maxDeviation {
		t.Errorf("Expected latency around %v, but got %v", expected, latency)
	}
}
