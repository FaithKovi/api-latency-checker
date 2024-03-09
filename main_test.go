package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetEndpointsFromURL(t *testing.T) {
	// Mock HTTP server to simulate responses
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<a href="http://example.com/page1"></a>`)
		fmt.Fprintln(w, `<a href="http://example.com/page2"></a>`)
	}))
	defer ts.Close()

	endpoints, err := getEndpointsFromURL(ts.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []string{"http://example.com/page1", "http://example.com/page2"}
	for _, endpoint := range expected {
		found := false
		for _, e := range endpoints {
			if e == endpoint {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Endpoint %s not found", endpoint)
		}
	}
}

func TestCheckLatency(t *testing.T) {
	// Mock HTTP server to simulate responses
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond) // Simulate delay
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	timeout := 1 * time.Second
	start := time.Now()
	latency, err := checkLatency(ts.URL, timeout)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if latency < 100*time.Millisecond {
		t.Errorf("Unexpected latency: %v", latency)
	}

	if elapsed > timeout {
		t.Errorf("Request took longer than expected")
	}
}
