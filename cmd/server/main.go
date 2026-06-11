package main

import (
	"log"
	"net/http"
	"time"
	
	"sse-demo/internal/broker"
	"sse-demo/internal/generators"
	"sse-demo/internal/handlers"
)

func main() {
	// Initialisation du broker
	b := broker.New()
	
	// Initialisation des générateurs de données
	stockGen := generators.NewStockGenerator(b)
	notifGen := generators.NewNotificationGenerator(b)
	userGen := generators.NewUserGenerator(b)
	
	// Démarrage des générateurs
	stockGen.Start(2 * time.Second)
	notifGen.Start(5 * time.Second)
	userGen.Start(7 * time.Second)
	
	// Initialisation des handlers
	sseHandler := handlers.NewSSEHandler(b)
	metricsHandler := handlers.NewMetricsHandler(b)
	
	// Configuration des routes
	http.HandleFunc("/events", sseHandler.Handle)
	http.HandleFunc("/clients", metricsHandler.HandleClients)
	http.Handle("/", http.FileServer(http.Dir(".")))
	
	// Démarrage du serveur
	log.Println("Serveur SSE démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}