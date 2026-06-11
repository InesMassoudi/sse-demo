package handlers

import (
	"fmt"
	"net/http"
	"time"
	
	"sse-demo/internal/broker"
	"sse-demo/pkg/sse"
)

// SSEHandler gère les connexions SSE
type SSEHandler struct {
	broker *broker.Broker
}

// NewSSEHandler crée un nouveau handler SSE
func NewSSEHandler(b *broker.Broker) *SSEHandler {
	return &SSEHandler{
		broker: b,
	}
}

// Handle gère les requêtes SSE
func (h *SSEHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Vérifier le support du streaming
	sseWriter, err := sse.NewWriter(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Configurer les en-têtes
	sse.SetHeaders(w)
	
	// Créer et enregistrer le client
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano())
	client := h.broker.Register(clientID)
	defer h.broker.Unregister(clientID)
	
	// Envoyer un événement de connexion
	_ = sseWriter.Send("connected", map[string]string{
		"clientId": clientID,
		"message":  "Connexion SSE établie !",
	})
	
	// Boucle principale
	for {
		select {
		case msg, ok := <-client.Channel:
			if !ok {
				return
			}
			if err := sseWriter.Send(msg.Event, msg.Data); err != nil {
				return
			}
		case <-r.Context().Done():
			// Client déconnecté
			return
		}
	}
}