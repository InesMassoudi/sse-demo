package generators

import (
	"math/rand"
	"time"
	
	"sse-demo/internal/broker"
	"sse-demo/internal/models"
)

// UserGenerator génère des activités utilisateur aléatoires
type UserGenerator struct {
	broker *broker.Broker
	names  []string
}

// NewUserGenerator crée un nouveau générateur d'activités utilisateur
func NewUserGenerator(b *broker.Broker) *UserGenerator {
	return &UserGenerator{
		broker: b,
		names:  []string{"Alice", "Bob", "Chaima", "David", "Emna", "Farid"},
	}
}

// Start démarre la génération des activités utilisateur
func (g *UserGenerator) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			g.generateActivity()
		}
	}()
}

func (g *UserGenerator) generateActivity() {
	name := g.names[rand.Intn(len(g.names))]
	action := "joined"
	if rand.Float32() < 0.4 {
		action = "left"
	}
	
	g.broker.Broadcast(models.Message{
		Event: "user_activity",
		Data: models.UserActivity{
			Name:   name,
			Action: action,
		},
	})
}