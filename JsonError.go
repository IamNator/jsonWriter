package jsonWrite

import (
	"encoding/json"
	"net/http"
)

func JsonError(w *http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	(*w).WriteHeader(ErrorCode)
	json.NewEncoder(*w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: ErrorMessage,
	})
}
