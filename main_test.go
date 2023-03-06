package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestWebserver(t *testing.T) {
	// Create a new Chi router
	r := chi.NewRouter()
	// Register a handler function for the root path
	r.Get("/", handlerFunc)

	// Create a new HTTP request for the root path
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to capture the server's response
	rr := httptest.NewRecorder()

	// Call the Chi router's ServeHTTP method to handle the request and capture the response
	r.ServeHTTP(rr, req)

	// Check that the server returns a 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that the response body is the expected string ("Hello, World!")
	expected := "Hello, World!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
