package main

import (
	"bytes"
	"context"
	"net/http"
	"testing"
)

// BenchmarkRESTGet - Test the REST POST connection
func BenchmarkRESTGet(b *testing.B) {
	client := &http.Client{}
	// Get to Posting
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "https://localhost:3001/12345", nil)
		req = req.WithContext(context.Background())
		client.Do(req)
	}
}

// BenchmarkRESTPost - Test the REST POST connection
func BenchmarkRESTPost(b *testing.B) {
	client := &http.Client{}

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
		req, _ := http.NewRequest(http.MethodPost, "https://localhost:3001", buf)
		req = req.WithContext(context.Background())
		client.Do(req)
	}
}

// BenchmarkRESTPut - Test the REST POST connection
func BenchmarkRESTPut(b *testing.B) {
	client := &http.Client{}

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
		req, _ := http.NewRequest(http.MethodPut, "https://localhost:3001/12345", buf)
		req = req.WithContext(context.Background())
		client.Do(req)
	}
}

// BenchmarkRESTDel - Test the REST POST connection
func BenchmarkRESTDel(b *testing.B) {
	client := &http.Client{}

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
		req, _ := http.NewRequest(http.MethodDelete, "https://localhost:3001/12345", buf)
		req = req.WithContext(context.Background())
		client.Do(req)
	}
}
