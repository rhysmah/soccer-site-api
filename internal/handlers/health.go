package handlers

import (
	"encoding/json"
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
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache, no-store")

		response := HealthResponse{
			Status:    "Healthy",
			Timestamp: time.Now(),
			Version:   server.GetVersion(),
			Uptime:    server.GetUptime().String(),
		}

		json.NewEncoder(w).Encode(response)
	}
}

type ServerProvider interface {
	GetVersion() string
	GetUptime() time.Duration
}
