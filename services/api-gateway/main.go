package main

import (
	"log"
	"net/http"
)

var httpPort = ":8080"

func main() {
	log.Println("API Gateway is starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/trip/details", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the API Gateway"))
	})

	server := &http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	log.Println("API Gateway is running on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
