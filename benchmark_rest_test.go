package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"net/http"
	"testing"
	"time"
)

var (
	restAddr   = "https://localhost:3001"
	restClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(3) * time.Millisecond,
	}
)

// BenchmarkRESTGet - Test the REST GET connection
func BenchmarkRESTGet(b *testing.B) {

	// Get to Getting
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, restAddr+"/vehicle/12345", nil)
		req = req.WithContext(context.Background())
		restClient.Do(req)
	}
}

// BenchmarkRESTPost - Test the REST POST connection
func BenchmarkRESTPost(b *testing.B) {

	// Create JSON Buffer
	buf := bytes.NewBufferString(`
		{
			"manufacturer":"Audi",
			"model":"A4",
			"price":24999.99,
			"mileage": 50000
		}
	`)

	// Get to Posting
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodPost, restAddr+"/vehicle", buf)
		req = req.WithContext(context.Background())
		restClient.Do(req)
	}
}

// BenchmarkRESTPut - Test the REST PUT connection
func BenchmarkRESTPut(b *testing.B) {

	// Create JSON Buffer
	buf := bytes.NewBufferString(`
		{
			"manufacturer":"Audi",
			"model":"A4",
			"price":24999.99,
			"mileage": 50000
		}
	`)

	// Get to Puting
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodPut, restAddr+"/vehicle/12345", buf)
		req = req.WithContext(context.Background())
		restClient.Do(req)
	}
}

// BenchmarkRESTDel - Test the REST DELETE connection
func BenchmarkRESTDel(b *testing.B) {

	// Get to Deleting
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodDelete, restAddr+"/vehicle/12345", nil)
		req = req.WithContext(context.Background())
		restClient.Do(req)
	}
}
