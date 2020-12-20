package jsWrite

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, Title string, Message string, StatusCode int) {
	(w).WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(struct {
		Title string  `json:"title"`
		Message string `json:"message"`
	}{
		Title: Title,
		Message: Message,
	})
}
