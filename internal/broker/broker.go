package broker

import (
	"log"
	"sync"
	"sse-demo/internal/models"
)

// Broker gère tous les clients connectés
type Broker struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

// New crée un nouveau broker
func New() *Broker {
	return &Broker{
		clients: make(map[string]*Client),
	}
}

// Register ajoute un nouveau client
func (b *Broker) Register(clientID string) *Client {
	client := NewClient(clientID)
	
	b.mu.Lock()
	b.clients[clientID] = client
	b.mu.Unlock()
	
	log.Printf("Client connecté : %s (total: %d)", clientID, b.Count())
	return client
}

// Unregister supprime un client
func (b *Broker) Unregister(clientID string) {
	b.mu.Lock()
	if client, ok := b.clients[clientID]; ok {
		client.Close()
		delete(b.clients, clientID)
	}
	b.mu.Unlock()
	
	log.Printf("Client déconnecté : %s (total: %d)", clientID, b.Count())
}

// Broadcast envoie un message à tous les clients
func (b *Broker) Broadcast(msg models.Message) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	for _, client := range b.clients {
		select {
		case client.Channel <- msg:
		default:
			// Canal plein, on ignore
			log.Printf("Canal plein pour le client %s", client.ID)
		}
	}
}

// Count retourne le nombre de clients connectés
func (b *Broker) Count() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.clients)
}

// GetClients retourne la liste des clients connectés
func (b *Broker) GetClients() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	clients := make([]string, 0, len(b.clients))
	for id := range b.clients {
		clients = append(clients, id)
	}
	return clients
}