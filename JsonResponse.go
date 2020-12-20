package jsonWrite

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, Title string, Message string, StatusCode int) {
	w.WriteHeader(StatusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Title string  `json:"title"`
		Message string `json:"message"`
	}{
		Title: Title,
		Message: Message,
	})
}
