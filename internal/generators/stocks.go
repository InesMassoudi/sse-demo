package generators

import (
	"math/rand"
	"time"
	
	"sse-demo/internal/broker"
	"sse-demo/internal/models"
)

// StockGenerator génère des cours boursiers aléatoires
type StockGenerator struct {
	broker *broker.Broker
	stocks map[string]float64
}

// NewStockGenerator crée un nouveau générateur de stocks
func NewStockGenerator(b *broker.Broker) *StockGenerator {
	return &StockGenerator{
		broker: b,
		stocks: map[string]float64{
			"AAPL": 175.0,
			"GOOG": 140.0,
			"MSFT": 380.0,
			"AMZN": 185.0,
		},
	}
}

// Start démarre la génération des cours
func (g *StockGenerator) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			g.updateStocks()
		}
	}()
}

func (g *StockGenerator) updateStocks() {
	for symbol, price := range g.stocks {
		change := (rand.Float64()*4 - 2) // entre -2 et +2
		newPrice := price + change
		g.stocks[symbol] = newPrice
		
		g.broker.Broadcast(models.Message{
			Event: "stock",
			Data: models.StockPrice{
				Symbol: symbol,
				Price:  newPrice,
				Change: change,
			},
		})
	}
}