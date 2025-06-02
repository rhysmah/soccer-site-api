package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rhysmah/soccer-site-api/internal/handlers"
	"github.com/rhysmah/soccer-site-api/internal/server"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page of soccer teams API server.\n")
}

func main() {

	server := interal.NewServer()

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/health", handlers.HealthHandler(server))

	log.Print("Starting server on :8080")
	log.Printf("Server version: %s", server.GetVersion())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
