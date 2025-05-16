package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("api/health", healthHandler)
	fmt.Println("Starting server on :7080...")
	if err := http.ListenAndServe(":7080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
