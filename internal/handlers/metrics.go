package handlers

import (
	"encoding/json"
	"net/http"
	
	"sse-demo/internal/broker"
)

// MetricsHandler gère les métriques du serveur
type MetricsHandler struct {
	broker *broker.Broker
}

// NewMetricsHandler crée un nouveau handler de métriques
func NewMetricsHandler(b *broker.Broker) *MetricsHandler {
	return &MetricsHandler{
		broker: b,
	}
}

// HandleClients retourne le nombre de clients connectés
func (h *MetricsHandler) HandleClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	response := map[string]interface{}{
		"count":   h.broker.Count(),
		"clients": h.broker.GetClients(),
	}
	
	json.NewEncoder(w).Encode(response)
}