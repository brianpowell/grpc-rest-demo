package main

import (
	"bytes"
	"net/http"
	"testing"
)

// BenchmarkRESTGet - Test the REST POST connection
func BenchmarkRESTGet(b *testing.B) {
	// Get to Posting
	for i := 0; i < b.N; i++ {
		http.Get("https://localhost:3001/12345")
	}
}

// BenchmarkRESTPost - Test the REST POST connection
func BenchmarkRESTPost(b *testing.B) {
	client := &http.Client{
		Transport: &http.Transport{},
	}

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
		client.Do(req)
	}
}

// BenchmarkRESTPut - Test the REST POST connection
func BenchmarkRESTPut(b *testing.B) {
	client := &http.Client{
		Transport: &http.Transport{},
	}

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
		client.Do(req)
	}
}

// BenchmarkRESTDel - Test the REST POST connection
func BenchmarkRESTDel(b *testing.B) {
	client := &http.Client{
		Transport: &http.Transport{},
	}

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
		client.Do(req)
	}
}
