package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Uptime    string    `json:"uptime"`
}

func HealthHandler(server ServerProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Status:    "Healthy",
			Timestamp: time.Now(),
			Version:   server.GetVersion(),
			Uptime:    server.GetUptime().String(),
		}

		// Encode to JSON in memory first, before touching HTTP response
		// Cannot alter HTTP response once written, but encoding may fail
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Printf("Failed to marshal health response: %v", err)
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)

			// Log any errors
			_, err := w.Write([]byte("Internal Server Error: failed to generate health check"))
			if err != nil {
				log.Printf("Failed to write health response: %v", err)
			}
			return
		}

		// JSON encoding succeeded
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache, no-store")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write(jsonData); err != nil {
			log.Printf("Failed to write health response: %v", err)
		}
	}
}

type ServerProvider interface {
	GetVersion() string
	GetUptime() time.Duration
}
