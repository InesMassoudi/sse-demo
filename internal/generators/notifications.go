package generators

import (
	"time"
	
	"sse-demo/internal/broker"
	"sse-demo/internal/models"
)

// NotificationGenerator génère des notifications périodiques
type NotificationGenerator struct {
	broker        *broker.Broker
	notifications []models.Notification
}

// NewNotificationGenerator crée un nouveau générateur de notifications
func NewNotificationGenerator(b *broker.Broker) *NotificationGenerator {
	return &NotificationGenerator{
		broker: b,
		notifications: []models.Notification{
			{"Nouveau message", "Ahmed vous a envoyé un message", "info"},
			{"Paiement reçu", "Transaction de 250€ confirmée", "success"},
			{"Alerte système", "CPU > 80% sur le serveur prod", "warning"},
			{"Déploiement réussi", "v2.3.1 est en production", "success"},
			{"Rapport prêt", "Le rapport mensuel est disponible", "info"},
		},
	}
}

// Start démarre la génération des notifications
func (g *NotificationGenerator) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		i := 0
		for range ticker.C {
			notification := g.notifications[i%len(g.notifications)]
			g.broker.Broadcast(models.Message{
				Event: "notification",
				Data:  notification,
			})
			i++
		}
	}()
}