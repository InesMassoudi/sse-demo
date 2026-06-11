package broker

import (
	"sse-demo/internal/models"
)

// Client représente un client connecté
type Client struct {
	ID      string
	Channel chan models.Message
}

func NewClient(id string) *Client {
	return &Client{
		ID:      id,
		Channel: make(chan models.Message, 10),
	}
}

// Close ferme le canal du client
func (c *Client) Close() {
	close(c.Channel)
}