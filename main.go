package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Path not found: %s\n", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()

	// Register the known route
	mux.HandleFunc("/api/health", healthHandler)

	// Fallback handler for all other routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only call notFoundHandler if the path doesn't match any existing handler
		// This check avoids overriding valid paths like /api/health
		_, pattern := mux.Handler(r)
		if pattern == "" {
			notFoundHandler(w, r)
		} else {
			// This should never be reached due to HandleFunc("/", ...) being last
			http.NotFound(w, r)
		}
	})

	fmt.Println("Starting server on :7080...")
	if err := http.ListenAndServe(":7080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
