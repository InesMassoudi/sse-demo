# SSE Real-Time Dashboard - Go Implementation

A demonstration project for Server-Sent Events (SSE) using a Go server and an HTML client.
No external dependencies — only the Go standard library (net/http, sync, encoding/json).

## Features

- **Scalable client management** with concurrent connection handling  
- **Multiple data streams** (stock prices, notifications, user activities)
- **Clean architecture** with separation of concerns
- **Production-ready error handling** and connection management
- **Automatic client reconnection** support

## Architecture

├── cmd/server/ # Application entry point
├── internal/
│ ├── broker/ # Client connection management
│ ├── handlers/ # HTTP request handlers
│ ├── generators/ # Data stream generators
│ └── models/ # Domain models
└── pkg/sse/ # Reusable SSE package


## Technical Functioning

### Server-Side Architecture

#### SSE Broker

The **Broker** manages all connected clients with a clean, concurrent-safe design:

Client connected → Register() → buffered Go channel created
Client disconnected → Unregister() → channel cleanly closed
New data → Broadcast() → sent to all active channel


**Key implementation details:**
- Each client gets its own buffered `chan Message` (buffer size: 10)
- `sync.RWMutex` protects the client map from concurrent access
- Non-blocking broadcasts prevent slow clients from blocking the system
- Automatic cleanup prevents memory leaks

#### 📡 SSE Handler (`GET /events`)
Client → GET /events
↓
Vérification http.Flusher
(streaming supporté ?)
↓
En-têtes HTTP obligatoires :
• Content-Type: text/event-stream
• Cache-Control: no-cache
• Connection: keep-alive
• Access-Control-Allow-Origin: *
↓
Register() client
Envoi événement "connected"
↓
Boucle principale :
┌─────────────────────────┐
│ Écoute channel client │ → Envoi SSE
│ OU │
│ r.Context().Done() │ → Déconnexion
└─────────────────────────┘
↓
Unregister() automatique (defer)


**Connection management:**
- Client disconnection detected via `r.Context().Done()`
- No goroutine leaks - proper cleanup guaranteed
- Deferred unregister ensures cleanup even on panic

#### 📨 SSE Event Format

```bash
# Stock price update
event: stock
data: {"symbol":"AAPL","price":176.34,"change":1.34}

# System notification
event: notification  
data: {"title":"Paiement reçu","message":"Transaction de 250€","level":"success"}

# User activity
event: user_activity
data: {"name":"Alice","action":"joined"}

🔄 Data Generators (Background Goroutines)
Generator	        Event Type	  Frequency	        Data Example
startStockTicker	stock	      Every 2 seconds	Stock price + variation
startNotifications	notification  Every 5 seconds	System alerts
startConnectedUsers	user_activity Every 7 seconds	Join/leave events


## Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/sse-demo

# Navigate to project
cd sse-demo

# Run the server
go run cmd/server/main.go

# Open your browser
http://localhost:8080
