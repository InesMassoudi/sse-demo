package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Writer gère l'écriture des événements SSE
type Writer struct {
	w       http.ResponseWriter
	flusher http.Flusher
}

// NewWriter crée un nouveau writer SSE
func NewWriter(w http.ResponseWriter) (*Writer, error) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("streaming non supporté")
	}
	
	return &Writer{
		w:       w,
		flusher: flusher,
	}, nil
}

// Send envoie un événement SSE
func (sw *Writer) Send(event string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erreur de marshalling: %w", err)
	}
	
	_, err = fmt.Fprintf(sw.w, "event: %s\n", event)
	if err != nil {
		return err
	}
	
	_, err = fmt.Fprintf(sw.w, "data: %s\n\n", jsonData)
	if err != nil {
		return err
	}
	
	sw.flusher.Flush()
	return nil
}

// SetHeaders définit les en-têtes HTTP pour SSE
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}