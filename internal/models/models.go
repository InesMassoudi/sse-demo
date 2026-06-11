package models

// Message représente un événement SSE
type Message struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

// StockPrice représente un cours boursier
type StockPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
}

type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Level   string `json:"level"` // info, success, warning
}

type UserActivity struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}