package jsonWrite

import (
	"encoding/json"
	"net/http"
)

//Writes Error Message to ResponseWriter in the form a json
//Similar to http.Error(w ResponseWriter, error string, code int)
func Error(w http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	w.WriteHeader(ErrorCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: ErrorMessage,
	})
}

