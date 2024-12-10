package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Route for the home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "Hello, Docker! <3")
	})

	// Route for health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		health := struct {
			Status string `json:"status"`
		}{Status: "OK"}
		json.NewEncoder(w).Encode(health)
	})

	// Get port from environment or default to 8080
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	// Start the server
	fmt.Printf("Starting server on :%s...\n", httpPort)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// Simple implementation of an integer minimum
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
